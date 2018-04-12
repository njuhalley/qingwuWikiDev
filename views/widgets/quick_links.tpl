<!-- 模态框（Modal） -->
<div class="modal fade" id="quickLinksModal" tabindex="-1" role="dialog" aria-labelledby="qucikLinksModalLabel">
    <div class="modal-dialog" role="document">
        <form method="post" id="addDocumentForm" class="form-horizontal">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4 class="modal-title" id="myModalLabel">QuickLinks</h4>
            </div>
            <div class="modal-body">
                <div class="form-group"><!--group 0 ## add by Qingwu-->
                    <div class="col-sm-6">
                        <a href="/api/schedules/edit/" class="btn btn-primary active" role="button" style="width:100%">
                            编辑-清无记事
                        </a>
                    </div>
                    <div class="col-sm-6">
                        <a href="/book/schedules/setting" class="btn btn-primary active" role="button" style="width:100%">
                            设置-清无记事
                        </a>
                    </div>
                </div>
                <div class="form-group"><!--group 1 ## add by Qingwu-->
                    <div class="col-sm-6">
                        <a href="/api/research/edit/" class="btn btn-primary active" role="button" style="width:100%">
                            编辑-Research
                        </a>
                    </div>
                    <div class="col-sm-6">
                        <a href="/book/research/setting" class="btn btn-primary active" role="button" style="width:100%">
                            设置-Research
                        </a>
                    </div>
                </div>
                <div class="form-group"><!--group 2 ## add by Qingwu-->
                    <div class="col-sm-6">
                        <a href="/api/computer/edit/" class="btn btn-primary active" role="button" style="width:100%">
                            编辑-CS及编程相关
                        </a>
                    </div>
                    <div class="col-sm-6">
                        <a href="/book/computer/setting" class="btn btn-primary active" role="button" style="width:100%">
                            设置-CS及编程相关
                        </a>
                    </div>
                </div>
                <div class="form-group"><!--group 3 ## add by Qingwu-->
                    <div class="col-sm-6">
                        <a href="/api/encyclopedia/edit/" class="btn btn-primary active" role="button" style="width:100%">
                            编辑-清无百科
                        </a>
                    </div>
                    <div class="col-sm-6">
                        <a href="/book/encyclopedia/setting" class="btn btn-primary active" role="button" style="width:100%">
                            设置-清无百科
                        </a>
                    </div>
                </div>
                <div class="form-group"><!--group 4 ## add by Qingwu-->
                    <div class="col-sm-6">
                        <a href="/api/basics/edit/" class="btn btn-primary active" role="button" style="width:100%">
                            编辑-科学文化基础知识
                        </a>
                    </div>
                    <div class="col-sm-6">
                        <a href="/book/basics/setting" class="btn btn-primary active" role="button" style="width:100%">
                            设置-科学文化基础知识
                        </a>
                    </div>
                </div>
                <div class="form-group"><!--group 5 ## add by Qingwu-->
                    <div class="col-sm-6">
                        <a href="/api/books/edit/" class="btn btn-primary active" role="button" style="width:100%">
                            编辑-清无典藏书库
                        </a>
                    </div>
                    <div class="col-sm-6">
                        <a href="/book/books/setting" class="btn btn-primary active" role="button" style="width:100%">
                            设置-清无典藏书库
                        </a>
                    </div>
                </div>
            </div>
            <div class="modal-footer">
                <span id="add-error-message" class="error-message"></span>
                <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            </div>
        </div>
        </form>
    </div>
</div>
<script type="text/javascript">
    $("#btn_quick_links").bind("click",function(){
        $('#quickLinksModal').modal({backdrop: "static"});
        //$('#quickLinksModal').modal("show");
    });
</script>