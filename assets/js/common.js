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