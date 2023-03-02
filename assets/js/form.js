// ajaxFormSubmit
// @param url
// @param formData
// @param href
// @param method
// @param submitBtn
function ajaxFormSubmit(url, formData, href = "/home", method = "POST",btn = null){
    $.ajax({
        url: url,
        method: method,
        data: formData,
        success: function(res) {
            console.log(res)
            if (res.Code != 200) {
                ErrorMessage(res.Message);
                if (btn) {
                    btn.text("提交").attr("disabled", false).removeClass("layui-disabled");
                }
            } else {
                SuccessMessage("提交成功！");
                setTimeout(function(){
                    window.location.href = href
                },1000)
                // alert("您当前通行证:" + res.token.token)

                // // 跳转页面
                // window.location.href = "/ERP-WEB/index.html";
            }
        },
        error: function(res) {
            if (res.responseJSON.message) {
                ErrorMessage(res.responseJSON.message);
            } else {
                ErrorMessage("提交失败！");
            }
            if (btn) {
                btn.text("提交").attr("disabled", false).removeClass("layui-disabled");
            }
        }
    })
}