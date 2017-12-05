import React, { Component } from "react";
import "../node_modules/font-awesome/css/font-awesome.css";
import "../node_modules/materialize-css/dist/css/materialize.css";
import "../node_modules/materialize-css/dist/js/materialize.js";

class Form extends Component {
    render() {
        return (
            <div class="money-form">
                <form
                    class="col s12"
                    onSubmit={this.props.handleFunc.createTran}
                >
                    <div class="row">
                        <div class="input-field col s3">
                            <input
                                placeholder=""
                                id="date"
                                type="text"
                                class="validate"
                            />
                            <label for="date">Date</label>
                        </div>
                    </div>
                    <label>Expense Type</label>
                    <select class="browser-default" id="type">
                        <option value="Food">Food</option>
                        <option value="Transportation">Transportation</option>
                        <option value="Shopping">Shopping</option>
                        <option value="Travel">Travel</option>
                    </select>
                    <div class="row">
                        <div class="input-field col s3">
                            <input
                                placeholder=""
                                id="amount"
                                type="text"
                                class="validate"
                            />
                            <label for="amount">Amount</label>
                        </div>
                    </div>
                    <div class="row">
                        <div class="input-field col s3">
                            <input
                                placeholder=""
                                id="note"
                                type="text"
                                class="validate"
                            />
                            <label for="note">Note</label>
                        </div>
                    </div>

                    <button
                        type="submit"
                        className="waves-effect waves-light btn yellow"
                    >
                        Add
                    </button>
                </form>
            </div>
        );
    }
}

export default Form;
