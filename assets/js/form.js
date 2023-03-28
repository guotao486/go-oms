
// ajaxFormSubmit
// @param url
// @param formData
// @param href
// @param method
// @param submitBtn
function ajaxSubmit(url, formData, href = "/home", method = "POST",btn = null){
    $.ajax({
        url: url,
        method: method,
        data: formData,
        success: function(res) {
            if (res.code != 200) {
                ErrorMessage(res.message);
                ShowErrorMessage(res)
                if (btn) {
                    btn.text("提交").attr("disabled", false).removeClass("layui-disabled");
                }
            } else {
                SuccessMessage("提交成功！");
                if (href != null){
                    setTimeout(function(){
                        window.location.href = href
                    },1000)
                }
                // alert("您当前通行证:" + res.token.token)

                // // 跳转页面
                // window.location.href = "/ERP-WEB/index.html";
            }
        },
        error: function(res) {
            console.log(res.responseJSON)
            if (res.responseJSON) {
                ErrorMessage(res.responseJSON.message);
                ShowErrorMessage(res.responseJSON)
            } else {
                ErrorMessage("提交失败！");
            }
            if (btn) {
                btn.text("提交").attr("disabled", false).removeClass("layui-disabled");
            }
        }
    })
}

function ajaxHtml(url, formData){
    $.ajax({
        url: url,
        method: "GET",
        data: formData,
        success: function(res) {
            layer.open({
                title: '#',
                type: 1,
                area: ['80%','80%'],
                content: res
            });
        },
        error: function(res) {
            if (res.responseJSON) {
                ErrorMessage(res.responseJSON.message);
            } else {
                ErrorMessage("提交失败！");
            }
        }
    })
}