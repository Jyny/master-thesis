import abstractview from "./abstractview.js";

export default class extends abstractview {
    constructor(params) {
        super(params);
        this.setTitle("Register");
    }

    async getHtml() {
        return `
            <h1>Welcome to New Session</h1>
            <p>
                If you want to register as an Session Owner. <br/> Click the button below.
            </p>
            <div class="input-group mt-5 mb-3">
                <div class="input-group-prepend">
                    <span class="input-group-text">Session ID</span>
                </div>
                <input type="text" class="form-control" value="${this.params.session_id}" readonly>
            </div>
            <p>
                <a class="btn btn-lg btn-secondary" id="reg_btn">Register As Owner</a>
            </p>
        `;
    }

    async listener() {
        const el = document.getElementById("reg_btn");
        const url = "/v1/meeting/" + this.params.session_id + "/owner"
        el.addEventListener("click", function() {
            let xhttp = new XMLHttpRequest();
            xhttp.open("POST", url, true);
            xhttp.onreadystatechange = function() {
                if (this.readyState == 4){
                    location.reload()
                }
            };
            xhttp.send();
        });
    }
}