import welcome from "./views/welcome.js";
import register from "./views/register.js";
import challenge from "./views/challenge.js";

document.addEventListener("DOMContentLoaded", () => {
    router();
});

const router = async () => {
    const appstate = "appstate"
    const routes = {
        welcome : {view: welcome },
        register : {view: register },
        challenge : {view: challenge },
        unseal: {},
    };

    let state = docCookies.getItem(appstate)
    let parmas = {
        session_id : docCookies.getItem("session_id"),
        owner_id : docCookies.getItem("owner_id"),
        owner_key : docCookies.getItem("owner_key"),
        challenge : docCookies.getItem("challenge"),
        answer: docCookies.getItem("answer")
    }

    let view = new routes[state].view(parmas)
    document.querySelector("#app").innerHTML = await view.getHtml();
    view.listener()
};

const docCookies = {
    getItem: function (sKey) {
        return decodeURIComponent(document.cookie.replace(new RegExp("(?:(?:^|.*;)\\s*" + encodeURIComponent(sKey).replace(/[-.+*]/g, "\\$&") + "\\s*\\=\\s*([^;]*).*$)|^.*$"), "$1")) || null;
    },
    setItem: function (sKey, sValue, vEnd, sPath, sDomain, bSecure) {
        if (!sKey || /^(?:expires|max\-age|path|domain|secure)$/i.test(sKey)) { return false; }
        var sExpires = "";
        if (vEnd) {
        switch (vEnd.constructor) {
            case Number:
            sExpires = vEnd === Infinity ? "; expires=Fri, 31 Dec 9999 23:59:59 GMT" : "; max-age=" + vEnd;
            break;
            case String:
            sExpires = "; expires=" + vEnd;
            break;
            case Date:
            sExpires = "; expires=" + vEnd.toUTCString();
            break;
        }
        }
        document.cookie = encodeURIComponent(sKey) + "=" + encodeURIComponent(sValue) + sExpires + (sDomain ? "; domain=" + sDomain : "") + (sPath ? "; path=" + sPath : "") + (bSecure ? "; secure" : "");
        return true;
    },
    removeItem: function (sKey, sPath, sDomain) {
        if (!sKey || !this.hasItem(sKey)) { return false; }
        document.cookie = encodeURIComponent(sKey) + "=; expires=Thu, 01 Jan 1970 00:00:00 GMT" + ( sDomain ? "; domain=" + sDomain : "") + ( sPath ? "; path=" + sPath : "");
        return true;
    },
    hasItem: function (sKey) {
        return (new RegExp("(?:^|;\\s*)" + encodeURIComponent(sKey).replace(/[-.+*]/g, "\\$&") + "\\s*\\=")).test(document.cookie);
    },
    keys: /* optional method: you can safely remove it! */ function () {
        var aKeys = document.cookie.replace(/((?:^|\s*;)[^\=]+)(?=;|$)|^\s*|\s*(?:\=[^;]*)?(?:\1|$)/g, "").split(/\s*(?:\=[^;]*)?;\s*/);
        for (var nIdx = 0; nIdx < aKeys.length; nIdx++) { aKeys[nIdx] = decodeURIComponent(aKeys[nIdx]); }
        return aKeys;
    }
};