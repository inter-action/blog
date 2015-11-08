#工作日志: 14-12-24 #JS 分页实现

这两天写了个JS分页的实现,贴在这里。

这个分页的设计思维就是把分页的视图层和分页的逻辑模型层分开。以便实现尽可能的组件复用。
逻辑模型做的工作就是渲染结果但是不关联样式和事件的绑定。分页的视图层需要将特定的样式传递给逻辑层获得渲染后的结果。绑定事件。


好, 个人表述能力就这么渣了，上代码吧

分页的逻辑模型

```javascript

    var models = (function(){
        //分页的视图模型
        var Pagination = (function(){
            var Pagination = function(results, options){
                options = options || {};

                this.cur_page = 1;
                this.num_per_page = options.num_per_page || 10;
                this.total = 0;
                this.set_results(results || []);
            };

            Pagination.prototype.set_results = function(results){
                this.results = results;
                this.total = this.results.length;
                if (results.length !== 0){
                    this.max_page = Math.ceil(this.total/this.num_per_page);
                }else{
                    this.max_page = 1;
                }
            };


            Pagination.prototype.get_results = function(pageno){
                assert(pageno >= 1 && pageno <= this.max_page);

                var start_idx = (pageno - 1) * this.num_per_page;
                var end_idx = pageno * this.num_per_page > this.total ? this.total : pageno * this.num_per_page; 

                return this.results.slice(start_idx, end_idx);
            };

            //渲染分页页面的结果
            Pagination.prototype.render_results = function(tpL_options){
                var node_tpl = tpL_options.node_tpl;
                var row_tpl = tpL_options.row_tpl;
                var results = this.get_results(this.cur_page);
                var h = [];


                for (var _i = 0; _i < results.length; _i++){
                    h.push( this.render_row(results[_i], node_tpl, row_tpl) );
                }

                return h.join('');
            };


            Pagination.prototype.render_row = function(row, node_tpl, row_tpl){
                var h = [];

                for (var _j = 0; _j < row.length; _j++){
                    h.push( node_tpl.format(row[_j]) );
                }

                return row_tpl.format(h.join(''));
            };

            //渲染分页页面的页脚
            Pagination.prototype.render_page_no = function(tpL_options){
                var node_tpl = tpL_options.node_tpl;
                var pre_tpl = tpL_options.pre_tpl;
                var next_tpl = tpL_options.next_tpl;
                var node_ellipsis = tpL_options.node_ellipsis;
                var MAX_NODES = 10;
                var ELLIPSIS_NODES_COUNT = 1;
                var CUR_PAGE_NODE_OFFSET = 2;
                var h = [];
                var start_page_no;

                if (this.cur_page !== 1){
                    h.push( pre_tpl.format(this.cur_page-1) );
                }

                var _i, _j;

                if (this.max_page >= MAX_NODES){

                    //最后n-ELLIPSIS_NODES_COUNT-1个
                    if (this.cur_page >= this.max_page - (MAX_NODES - ELLIPSIS_NODES_COUNT - 1) + 1){
                        h.push( node_tpl.format(1));

                        for (_j = 0; _j < ELLIPSIS_NODES_COUNT; _j++ ){
                            h.push(node_ellipsis);
                        }

                        start_page_no = this.max_page - (MAX_NODES - ELLIPSIS_NODES_COUNT - 1) + 1;
                        for (_i = start_page_no; _i < this.max_page; _i++){
                            h.push( node_tpl.format( _i ) );
                        }
                    }else {

                        if (this.cur_page <= CUR_PAGE_NODE_OFFSET){//当前页左偏移范围内
                            for (_i = 0; _i < MAX_NODES - ELLIPSIS_NODES_COUNT - 1; _i++){
                                h.push( node_tpl.format( (_i+1) ) );
                            }
                        }else{
                            start_page_no = this.cur_page - CUR_PAGE_NODE_OFFSET;
                            for (_i = 0; _i < MAX_NODES - ELLIPSIS_NODES_COUNT - 1; _i++){
                                h.push( node_tpl.format(start_page_no + _i) );
                            }
                        }

                        for (_j = 0; _j < ELLIPSIS_NODES_COUNT; _j++ ){
                            h.push(node_ellipsis);
                        }
                    }

                    h.push( node_tpl.format(this.max_page) );//last node

                }else{
                    for (_i = 0; _i < this.max_page; _i++ ){
                        h.push(node_tpl.format(_i+1));
                    }
                }


                if (this.cur_page !== this.max_page){
                    h.push( next_tpl.format(this.cur_page + 1) );
                }


                return h.join('');
            };


            return Pagination;
        })();


        return {
            Pagination: Pagination
        };
    })();
```

分页的视图层(jquery 插件扩展), 用的是 amaze-ui 的分页样式

```javascript
(function($){
    $.fn.extend({
        /**
         * [pagination description]
         * @param  {[type]} pagemodel       分页的视图的模型
         * @param  {[type]} $results_anchor 结果页面的节点
         * @return {[type]}                 [description]
         */
        pagination: function(pagemodel, $results_anchor, callback){
            var _page = pagemodel;
            callback = callback || function(){};

            return this.each(function(){
                var $no_anchor = $(this);

                var page_tpl = '<ul data-am-widget="pagination" class="am-pagination am-pagination-default" style="width: 100%; text-align: center;">#{0}</ul>';
                var node_tpls = {
                    node_tpl: '<li class=""> <a href="javascript: void(0);" data-no="#{0}">#{0}</a> </li>',
                    pre_tpl:  '<li class="am-pagination-prev "> <a href="javascript: void(0);" data-no="#{0}">上一页</a> </li>',
                    next_tpl:  '<li class="am-pagination-next "> <a href="javascript: void(0);" data-no="#{0}">下一页</a> </li>',
                    node_ellipsis: '<li class=""> <a href="#">...</a> </li>'
                };

                function render_pagination(pageno){
                    _page.cur_page = pageno;


                    var results_html = _page.render_results({node_tpl: '<td>#{0}</td>', row_tpl: '<tr>#{0}</tr>'});
                    $results_anchor.empty().append(results_html);
                    

                    var no_html = page_tpl.format(_page.render_page_no(node_tpls));
                    var $no = $(no_html);
                    //当前页高亮
                    $no.find('li>a[data-no="#{0}"]'.format(_page.cur_page)).parent('li').addClass('am-active');
                    $no_anchor.empty().append($no);


                    $no_anchor.find('a[data-no]').on('click', function(){
                        var pageno = $(this).data('no');
                        render_pagination(pageno);
                    });
                }

                render_pagination(_page.cur_page);

                callback();
            });
        }
    });


})(jQuery);

```

用法

```html
    <div>
        <table class="am-table am-table-bordered am-table-radius am-table-striped">
            <thead id="thead">

            </thead>
            <tbody id="tbody">

            </tbody>
        </table>
    </div>


    <div id="pagination">
    </div>

```

//---


```javascript
var _page = new models.Pagination(results);
$('#pagination').pagination(_page, $('#tbody'), function(){
    $('#thead').html('<tr><th>日期</th> <th>票数</th> <th>购买金额</th> <th>中奖金额</th> </tr>');
});
```


//渲染的结果截图
![常用的分页](http://m2.img.srcdd.com/farm5/d/2014/1224/12/910A5DABB9AF0F876A0C3E71653C4D5E_B500_900_500_330.png)



好了,常用的逻辑模型就是上边那样，现在呢说下扩展
我现在有个需求，就是每两行是一组数据，一组数据在分页中才能看做一行，好，扩展下分页的逻辑模型

```javascript
//带彩种的分页数据模型
var LotteryIdsPagintionModel = (function(_super){
    function LotteryIdsPagintionModel(results, options){
        LotteryIdsPagintionModel.__super__.constructor.call(this, results, options);
    }
    UTIL.OOP.__extends(LotteryIdsPagintionModel, _super);
    
    /*
        @param this.results:

        [
            [
                [date, name, v1, v2, v3]
                [date, name, v1, v2, v3]
                [date, name, v1, v2, v3]
            ]
            ...
        ]
     */
    //@Override
    LotteryIdsPagintionModel.prototype.render_results = function(tpL_options){
        //var node_tpl = tpL_options.node_tpl;
        var row_tpl = tpL_options.row_tpl;
        var results = this.get_results(this.cur_page);
        var h = [];


        /*
        @param row
            [
                [date, name, v1, v2, v3]
                [date, name, v1, v2, v3]
            ]
         */
        function _render_row(row){
            var h = [], _t;

            for (var i = 0; i<row.length; i++){
                var _row = row[i];
                if (i===0){
                    _t = '<td rowspan="#{5}">#{0}</td><td>#{1}</td><td>#{2}</td><td>#{3}</td><td>#{4}</td>'.format(_row[0], _row[1], _row[2], _row[3], _row[4], row.length);
                    h.push( row_tpl.format(_t) );//todo: paramterize thisrow
                }else{
                    _t = '<td>#{0}</td><td>#{1}</td><td>#{2}</td><td>#{3}</td>'.format(_row[1], _row[2], _row[3], _row[4]);
                    h.push( row_tpl.format(_t) );
                }
            }
            return h.join('');
        }

        for (var _i = 0; _i < results.length; _i++){
            h.push( _render_row(results[_i]) );
        }

        return h.join('');
    };

    return LotteryIdsPagintionModel;
})(models.Pagination);
```

用法

```javascript
//这里视图层只需要更改逻辑模型就能实现所要定制的功能了
var _page = new LotteryIdsPagintionModel(results);
$('#pagination').pagination(_page, $('#tbody'), function(){
    $('#thead').html('<tr><th>日期</th><th>彩种</th> <th>票数</th> <th>购买金额</th> <th>中奖金额</th> </tr>');
});
```


//渲染后的结果截图

![n行一组的分页](http://m2.img.srcdd.com/farm4/d/2014/1224/12/A532ECBA38E322C7A59ECE2419D0F1FB_B500_900_500_396.png)