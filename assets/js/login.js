particlesJS.load('particles-js', '/assets/js/particles.json', function() {
    console.log('callback - particles.js config loaded');
});

layui.use(['form', 'layer'], function(){
    // 表单提交事件
    var form = layui.form;
    form.on('submit(form_submit)', function(data){
        formData = $('.layui-form').serialize()
        var btn = $(this);
        btn.text("提交中...").attr("disabled", "disabled").addClass("layui-disabled");
        $.ajax({
            url: "/user/create",
            method: "POST",
            data: formData,
            success: function(res) {
                if (res.code != 200) {
                    layer.msg(res.msg);
                    btn.text("提交").attr("disabled", false).removeClass("layui-disabled");
                } else {
                    //把token写到cookie
                    //登陆成功这后把token放到cookie里面
                    // $.cookie('TOKEN', res.token.token, {
                    //     expires: 7
                    // });

                    // // 将权限存入到浏览器
                    // localStorage.setItem("permissions", res.token.permissions);
                    // // 将用户类型存入到浏览器
                    // localStorage.setItem("usertype", res.token.usertype);
                    // // 将用户名存入浏览器
                    // localStorage.setItem("username", res.token.username);

                    // alert("您当前通行证:" + res.token.token)

                    // // 跳转页面
                    // window.location.href = "/ERP-WEB/index.html";
                }
            },
            error: function(res) {
                console.log(res)
                layer.msg("");
                btn.text("提交").attr("disabled", "").removeClass("layui-disabled");
            }
        })
        return false;
    });
})