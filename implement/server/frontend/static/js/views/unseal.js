import abstractview from "./abstractview.js";

export default class extends abstractview {
    constructor(params) {
        super(params);
        this.setTitle("Seesion Owner");
    }

    async getHtml() {
        return `
            <h1>Unseal to Record</h1>
            <p>
                Challenge Succeeded ! <br/>
                If every Owner Challenge Succeeded, <br/>
                Click the <code>Unseal</code>  button below to nseal. <br/>
                After unseal completed, click <code>Access Record</code> to Download Record.
            </p>
            <div class="input-group mt-5 mb-3">
                <div class="input-group-prepend">
                    <span class="input-group-text">Session ID</span>
                </div>
                <input type="text" class="form-control" value="${this.params.session_id}" readonly>
            </div>
            <div class="input-group mb-3">
                <div class="input-group-prepend">
                    <span class="input-group-text">Owner ID</span>
                </div>
                <input type="text" class="form-control" value="${this.params.owner_id}" readonly>
            </div>
            <p>
                <a class="btn btn-lg btn-secondary" id="unseal_btn">Unseal</a>
            </p>
            <p>
                <a class="btn btn-lg btn-secondary" id="dl_btn">Access Record</a>
            </p>
        `;
    }

    async listener() {
        const el = document.getElementById("unseal_btn");
        const url = "/v1/unseal/" + this.params.session_id
        const ownerkey = this.params.owner_key
        el.addEventListener("click", function() {
            let xhttp = new XMLHttpRequest();
            xhttp.open("GET", url, true);
            xhttp.onreadystatechange = function() {
                if (this.readyState == 4){
                    location.reload()
                }
            };
            xhttp.send();
        });
    }
}