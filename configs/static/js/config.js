/*=========================================公共变量信息-开始=========================================*/
let publicUrl = "/data/"
//列表图片url地址前缀
let prefix_url = "https://img.kuakedata.com/"
let login_url = "/login"
let url_404 = "/404"
let url_500 = "/500"
let expire_date = (new Date().getTime() / 1000) + 86400
/*=========================================基础数据变量信息-开始=========================================*/
//发送验证码
let url_verify = '/common/verify'
let url_login = '/common/login'
let url_get_sub_list = '/subject/get/list'
let url_set_sub_try = '/subject/set/try'
let url_get_sub_detail = '/subject/get/detail'
let url_get_cha_list = '/chapter/get/list'

/*=========================================公共方法-开始=========================================*/

//获取请求地址中携带的参数
function getRequestParam() {
    let obj = {};
    let arr = window.location.search.slice(1).split("&");
    for (let i = 0, len = arr.length; i < len; i++) {
        let nv = arr[i].split("=");
        obj[decodeURIComponent(nv[0]).toLowerCase()] = decodeURIComponent(nv[1]);
    }
    return obj;
}

//全局加载提示
function show_loading() {
    this.loadingMsg = layer.msg('加载中，请稍后...',
        {
            /*icon: 16,*/
            id: 'layer_loading' + new Date().getTime(),
            zIndex: layer.zIndex,
            offset: '100px',
            shade: [0.1],
            scrollbar: false,
            time: 0,
            success: function (layero) {
                layer.setTop(layero);
            }
        });
}

function hide_loading() {
    layer.close(this.loadingMsg);
}

function sendApi(url, data, sendType, token) {
    if (token == undefined) {
        token = ""
    }
    let result = {};
    $.ajax({
        type: sendType,
        url: url,
        dataType: 'json',
        data: data,
        headers: {"token": token},
        async: false,
        beforeSend: function () {
        },
        success: function (res) {
            result = res
            jump_login(result)
        },
        complete: function () {
            hide_loading()
        },
        error: function (res) {
            hide_loading()
            result = res
            layer.msg("服务器繁忙！接口返回空数据", {anim: 6});
            return false
        }
    })
    return result
}

function jump_login(result) {
    let hash = parent.location.hash;
    if (result.code == 30008 || result.code == 20002) {
        layer.msg('您还没登录，请先去登录', {shade: 0.5, time: 3000}, function () {
            window.parent.frames.location.href = login_url + hash
        });
    } else if (result.code != 0 && result.code != 30008 && result.code != 20002) {
        layer.msg(result.msg, {anim: 6});
    } else if (result.code == undefined && result.code != 30008 && result.code != 20002) {
        layer.msg("服务器繁忙！接口返回空数据", {anim: 6});
    }
    return false
}

/*=========================================公共方法-结束=========================================*/

