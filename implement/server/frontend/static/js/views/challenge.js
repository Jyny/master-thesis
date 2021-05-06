import abstractview from "./abstractview.js";

export default class extends abstractview {
    constructor(params) {
        super(params);
        this.setTitle("Seesion Owner");
    }

    async getHtml() {
        return `
            <h1>Unseal Challenge</h1>
            <p>
                Solve the Challengeto Unseal the Record. <br/>
                Decrypt challenge to <code>Solve and Sign</code> with Owner Key.
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
            <div class="accordion mb-3" id="accordionExample">
                <div class="accordion-item">
                    <h2 class="accordion-header" id="headingOne">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseOne" aria-expanded="flase" aria-controls="collapseOne">
                        Owner Key
                    </button>
                    </h2>
                    <div id="collapseOne" class="accordion-collapse collapse" aria-labelledby="headingOne" data-bs-parent="#accordionExample">
                        <div class="accordion-body">
                            <code>${this.params.owner_key}</code>
                        </div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="headingTwo">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseTwo" aria-expanded="false" aria-controls="collapseTwo">
                        Challenge
                    </button>
                    </h2>
                    <div id="collapseTwo" class="accordion-collapse collapse" aria-labelledby="headingTwo" data-bs-parent="#accordionExample">
                        <div class="accordion-body">
                            <code>${this.params.challenge}</code>
                        </div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="headingThree">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseThree" aria-expanded="false" aria-controls="collapseThree">
                        Solve
                    </button>
                    </h2>
                    <div id="collapseThree" class="accordion-collapse collapse" aria-labelledby="headingThree" data-bs-parent="#accordionExample">
                        <div class="accordion-body">
                            <code>${this.params.answer}</code>
                        </div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="headingThree">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseThree" aria-expanded="false" aria-controls="collapseThree">
                        Sign
                    </button>
                    </h2>
                    <div id="collapseThree" class="accordion-collapse collapse" aria-labelledby="headingThree" data-bs-parent="#accordionExample">
                        <div class="accordion-body">
                            <code>${this.params.sign}</code>
                        </div>
                    </div>
                </div>
            </div>
            <p>
                <a class="btn btn-lg btn-secondary" id="solve_btn">Solve Challenge</a>
            </p>
        `;
    }

    async listener() {
        const el = document.getElementById("solve_btn");
        const url = "/v1/unseal/challenge/" + this.params.session_id + "/" + this.params.owner_id
        const solve = this.params.answer
        const sign = this.params.sign
        el.addEventListener("click", function() {
            let xhttp = new XMLHttpRequest();
            xhttp.open("PUT", url, true);
            xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
            xhttp.onreadystatechange = function() {
                if (this.readyState == 4){
                    location.reload()
                }
            };
            xhttp.send(JSON.stringify({
                "solve" : solve,
                "sign": sign
            }));
        });
    }
}