前一段时间阅读了一篇酷壳上的文章 [程序的本质复杂性和元语言抽象](http://coolshell.cn/articles/10652.html#more-10652) 让我对元编程这个陌生又好奇的概念有了一个初步的了解。
加上最近公司要做一个JS的表单验证, 我呢，不得不说很懒，不喜欢用别人写的代码 copy-paste-change。所以我就想写一个通用的工具类，在做表单验证的时候，只需要输入规则就行，抽离的验证的逻辑。进一步抽象代码。
在这里不得不说前面读的 AngularJS 的一些风格影响到我。

代码由于没有进行任何的单元测试, 这里只进行参考交流用,并不做任何保证。

### 定义FormValidator
这里呢，这个类的依赖已经列出来了, util.js是我自己写的工具类。这个FormValidator用到的只有一个String.format和 [tiny sub/pub 的事件分发库](https://gist.github.com/cowboy/661855)

这里用到了JS的 modular pattern。没别的作用就是方便我自己的代码折叠。

```javascript

// -------------------- START: FormValidator --------------------------
// dependencies, underscore.js, jquery.js, util.js
// 
// this validator will publish a `'/form_validator/failure/<form_name>/<field_name>` failure msg
// with params [fieldname, strategy_name]
// 
// this validator donesnt support real-time validation yet
// 
// @param <form> :required
// @param [override_strategy] :optional
// 
// noempty:243,2,43:|regex|num|length|checked#blur

var FormValidator = (function(){
    
    // dependencies, underscore.js, jquery.js, util.js
    // 
    // this validator will publish a `'/form_validator/failure/<form_name>/<field_name>` failure msg
    // with params [fieldname, strategy_name]
    // 
    // this validator donesnt support real-time validation yet
    // 
    // @param <form> :required
    // @param [override_strategy] :optional
    // 
    // noempty:243,2,43:|regex|num|length|checked|
    function FormValidator(form, override_strategy) {
        this.$form = $(form);
        this.form_name = this.$form.attr('name');
        this.failure_msg_domain = '/form_validator/failure/' + this.form_name;
        this.success_msg_domain = '/form_validator/success/' + this.form_name;

        // should return true or false to indicate validation result
        var strategies = {
            'noempty': function() {
                var value = this.value.trim();

                if (!value) {
                    return false;
                }

                return true;
            },
            //regex:\\d{2}
            'regex': function() {
                var value = this.value;
                var regex_str = this.getAttribute('data-regex');

                assert(regex_str);

                var reg = new RegExp(regex_str, 'g');

                if (!reg.test(value)) {
                    return false;
                }
                return true;
            },
            //todo: 这个还没有测试过
            // number range validator, range include
            'number': function(min, max) {
                var value = this.value;
                var num_reg = /^\d+$/g;

                if (!num_reg.test(value)) {
                    return false;
                }

                min = min || Number.MIN_VALUE;
                max = max || Number.MAX_VALUE;

                value = parseInt(value, 10);

                return (value >= min && value <= max);
            },
            // length range validator, included
            'length': function(min, max) {
                var length = this.value.length;
                return (length >= min && length <= max);
            },
            //checkstatus:0|1
            'checkstatus': function(status) {
                if (status == null) {//check null and undifined
                    status = true;
                }

                if (status === '0') {
                    status = false;
                } else {
                    status = true;
                }

                return $(this).prop('checked') === status;

            }
        };

        this.strategies = $.extend({}, strategies, override_strategy);

        this.bind_events();
    }
    FormValidator.prototype.bind_events = function(){
        var fields = this.$form.find(':input[name]');
        var ctx = this;

        function _f(){
            ctx.validate(this);
        }

        for (var _i = 0; _i < fields.length; _i++) {
            var field = fields[_i];
            var name = field.name;
            var strategies = field.getAttribute('data-strategies');

            if (!strategies) {
                console.log('input name=#{0} has no strategies specified'.format(name));
                continue;
            }

            var _t  = this.parse_strategies(strategies);
            var es = _t[1];

            for (var _k = 0; _k < es.length; _k++){
                var e = es[_k];
                $(field).on(e, _f);
            }
        }

    };

    FormValidator.prototype.do_validate = function() {
        var fields = this.$form.find(':input[name]');
        var result = true;

        for (var i = 0; i < fields.length; i++) {
            result = this.validate(fields[i]);

            if (!result) {
                break;
            }
        }

        return result;
    };


    FormValidator.prototype.validate = function(input_ele) {
        var strategies = input_ele.getAttribute('data-strategies');
        var error_seq = input_ele.getAttribute('data-error_seq');//保证事件绑定的时候的顺序验证
        error_seq = error_seq ? parseInt(error_seq, 10) : error_seq;

        var name = input_ele.name;

        if (!strategies) {
            console.log('input name=#{0} has no strategies specified'.format(name));
            return true;
        }

        var ts = this.parse_strategies(strategies);
        var tokens = ts[0];

        var value = this.value;

        var result = true;

        for (var i = 0; i < tokens.length; i++) {
            var token = tokens[i];

            if (!error_seq || error_seq<=i){//顺序控制，当前面的strategries 验证失败的时候，只对失败的重新校验
                result = this.apply_strategy(token, value, input_ele);
            }

            if (!result) {
                input_ele.setAttribute('data-error_seq', i);
                break;
            }else {
                if (i === error_seq){//验证通过,移除失败的flag
                    input_ele.removeAttribute('data-error_seq', i);
                }
            }
        }

        return result;
    };

    FormValidator.prototype.apply_strategy = function(token, value, input_ele) {
        assert(token);
        assert(input_ele);

        var name = input_ele.name;
        assert(name);


        var _t = this.parse_token(token);
        var strategy_name = _t[0];
        var strategy_params = _t[1];

        var result = this.strategies[strategy_name].apply(input_ele, strategy_params);
        if (!result) { //validation failed
            $.publish(this.failure_msg_domain, [name, strategy_name]);
            this.$form.data('validation_failure_field', name);
        }else{
            $.publish(this.success_msg_domain, [name, strategy_name]);
        }

        return result;
    };


    // parse token_str to [strategy_name, <strategy_params>]
    //@param token_str 
    //          pattern:
    //              a:1,2,3
    FormValidator.prototype.parse_token = function(token_str) {
        var _t = token_str.split(':');
        var _p = [];
        if (_t.length >= 2) {
            _p = _.map(_t[1].split(','), function(str) {
                return str.trim();
            });
        }
        return [_t[0], _p];
    };
    // strategries_pattern:
    //      <stragtegy_token>|[stragtegy_token..]#[event,..]
    FormValidator.prototype.parse_strategies = function(strategies_str){
        var ts = strategies_str.split('#');
        assert(ts.length <= 2);

        var _es = [];
        if (ts.length === 2){
            _es = _.map(ts[1].split(','), function(event_str){
                return event_str.trim();
            });    
        }

        return [ts[0].split('|'), _es];
    };

    return FormValidator;
})();
// -------------------- END: FormValidator ----------------------------
```

好下面呢，我们来看下使用
所有需要验证的表单都需要name属性。需要验证的规则用 data-strategies 指定.

---

```plain

data-strategries的格式:  
    <strategy_ name>:[strategy_params]|[another_strategy...]#[bind_event_name]
[] 为选填, <> 为必填

```

---

```html

    <!-- 注册 -->
    <div class='alert_reg' id="pop_up_register" style="visibility: hidden; position: fixed;">
      <div class='x' data-action="close">
        <img src='/images/quickbuy/reg_h.png'/>
        <a href='javascript:void(0);'></a>
      </div>
      <form name="reg_form">
        <div class='reg_form'>
          <ul>
            <li id="passwdHG">
              <label class="usr">手机号：</label>
              <input type='text' name="phone" autocomplete="off" maxlength="11" data-strategies="regex|syn_phone_check#blur" data-regex="^1\d{10}$"/>
              <span data-id="info" data-rel="phone" data-msg-error="手机号码错误！"></span>
            </li>
            <li id="passwdHG">
              <label for="passwd">创建密码：</label> 
              <input type="password" name="password" data-strategies="regex#blur" data-regex="^[a-z0-9A-Z]{6,15}$" maxlength="15" />
              <span data-id="info" data-rel="password" data-msg-origin="6-15个字符(字母和数字)" data-msg-error="密码格式错误"></span>
            </li>
            <li id="confirPHG">
              <label for="confirmPasswd">确认密码：</label>
              <input type="password" name="repassword" data-strategies="identical:password#blur" maxlength="15" />
              <span data-id="info" data-rel="repassword" data-msg-origin="6-15个字符(字母和数字)" data-msg-error="两次密码不一致"></span>
            </li>
            <li id="authCodeHG">
              <label for="authCode">验证码：</label>
              <input type="text" name="authcode" autocomplete="off" data-strategies="length:4,4|captcha#blur" maxlength="4">
              <span id="authCodeShow"><img data-id="auth_img" width="52" height="22" align="absmiddle" style="cursor:hand;"title="点击刷新">&nbsp;</span>
              <img src="/images/newpic/ybz.png" data-id="refresh-img">
              <img data-id="info" data-rel="authcode" data-src-failure="/images/regist/error.jpg" data-src-success="/images/regist/true.jpg">
            </li>
          </ul>
          <div class='agree'>
            <div class='agree1'>
              <input type='checkbox' name="agechk" checked="true" data-strategies="checkstatus"/>
              <span>我已经年满18岁并同意</span>
              <a href='http://info.jifencai.com/html/help/service/'>《委托投注规则》</a>
              <span data-id="info" data-rel="agechk" data-msg-error="请阅读并同意接受 《委托投注规则》"></span>
            </div>
            <div class='agree2'>以下内容为选填项</div>
          </div>
          <ul>
            <li>
              <label>真实姓名：</label>
              <input type="text" name="realname" data-strategies="required_with:certno|regex" data-regex="^$|^[\u4e00-\u9fa5a-zA-z\.]+$" autocomplete="off" maxlength="20">
              <span data-id="info" data-rel="realname" data-msg-error="请输入有效的真实姓名">请输入真实姓名</span>
            </li>
            <li>
              <label>身份证号：</label>
              <input type="text" name="certno" data-strategies="required_with:realname|regex" data-regex="^$|^\d{15}$|^\d{18}$|^(\d{17}[Xx])$" autocomplete="off" maxlength="18">
              <span data-id="info" data-rel="certno" data-msg-error="请输入有效的身份证号">请输入有效的身份证号</span>
            </li>
          </ul>
        </div>
        <div class='tj_reg'>
          <input type="button" name="submit" value="提交注册" style="cursor: pointer;" data-action="submit">
        </div>
      </form>
    </div>

```


由于我的功能是需要在用户输入的时候验证, 提交的时候又要验证。所以呢,在上边的一些字段中绑定了blur事件时候触发。

好，下面我们说下视觉反馈是如何做的

视觉反馈是需要你自己处理的。所以上边的FormValidator是没有任何视觉处理的逻辑的。这个任务就交给了你自己。你需要监听 failure / success 两个事件。并处理相应的视觉提示逻辑就ok了。这里也展示了如何扩展你的验证逻辑。这里扩展了2次密码需要一致。


下面是实时校验的逻辑：


```javascript
REGISTER_CONTEXT.failure_msg_domain = '/form_validator/failure/#{0}'.format(_jquery_map.$form.attr('name'));
REGISTER_CONTEXT.success_msg_domain = '/form_validator/success/#{0}'.format(_jquery_map.$form.attr('name'));
REGISTER_CONTEXT.validator = new FormValidator(_jquery_map.$form[0], {
    identical: function(target_name) {
        var password = _jquery_map.$form.find(':password[name=#{0}]'.format(target_name)).val();
        var repassword = this.value;

        console.debug('validate password #{0}, repassword: #{1}'.format(password, repassword));
        if (password === repassword) {
            return true;
        }

        return false;
    },

    required_with: function(target_name) {
        var t_value = _jquery_map.$form.find(':text[name=#{0}]'.format(target_name)).val();
        var value = this.value;

        return !(t_value.trim() && !value.trim());
    }
});

$.subscribe(REGISTER_CONTEXT.failure_msg_domain, function($event, name, strategy_name) {
    console.warn('failed to validate feild: #{0}, with strategy: #{1}'.format(name, strategy_name));
    var span =  _jquery_map.$infos.filter('[data-rel=#{0}]'.format(name));
    var err_msg = span.attr('data-msg-error');
    if (name === 'agechk'){
        sAlert(err_msg);
        return ;
    }

    if (err_msg) {
        _jquery_map.$infos.filter('[data-rel=#{0}]'.format(name)).html(err_msg);
    }
});

$.subscribe(REGISTER_CONTEXT.success_msg_domain, function($event, name){
    var span = _jquery_map.$infos.filter('[data-rel=#{0}]'.format(name));
    var origin_msg = span.attr('data-msg-origin');

    if (origin_msg){
        span.html(origin_msg);
    }else{
        span.html('');
    }
});

```

表单提交前的校验
REGISTER_CONTEXT.validator.do_validate();这样一句就行了


这样这个类的用法我就讲完了。希望对你有些收益。