
function ErrorMessage(msg){
    layer.msg(msg, {icon: 2});
}

function SuccessMessage(msg){
    layer.msg(msg, {icon: 1});
}

function WarningMessage(msg){
    layer.msg(msg,{icon:0})
}

function Loading(){
    var index = layer.load(2, {
        shade:[0.8, '#393D49'],
        shadeClose:true
    })
    return index
}

function LayerClose(index){
    layer.close(index);
}

function HideAllMessage() {
    $("#error").hide()
    $("#warn").hide()
    $("#success").hide()
}

function HideMessage(e){
    e.parent().parent().hide()
}

function ShowMessage(res, type = "error") {
    that = $("#"+type)
    that.show()
    html = "<p>"+res.message+"<p>"
    if (res.details){
        for (var i=0;i<res.details.length;i++)
        { 
            html += "<p>" +res.details[i]+ "</p>"
        }
    }
    that.find('.message').html(html)
}

function ShowErrorMessage(res) {
    ShowMessage(res, "error")
}
function ShowWarnMessage(res) {
    ShowMessage(res, "warn")
}
function ShowSuccessMessage(res) {
    ShowMessage(res, "success")
}