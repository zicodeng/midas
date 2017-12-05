import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";
import Transaction from "./Transaction.js";
import "../node_modules/font-awesome/css/font-awesome.css";
import "../node_modules/materialize-css/dist/css/materialize.css";
import "../node_modules/materialize-css/dist/js/materialize.js";
import Form from "./Form.js";

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            transactions: [
                {
                    date: "11/27",
                    type: "Food",
                    amount: "200",
                    note: "Costco"
                },
                {
                    date: "11/22",
                    type: "Shopping",
                    amount: "30",
                    note: "American Eagle"
                },
                {
                    date: "11/21",
                    type: "Travel",
                    amount: "320",
                    note: "Flight to SFO"
                },
                {
                    date: "11/15",
                    type: "Travel",
                    amount: "100",
                    note: "Flight to Portland"
                },
                { date: "11/02", type: "Food", amount: "30", note: "Chicken" }
            ],
            adding: false,
            type: "all"
        };
    }

    createTran(event) {
        event.preventDefault();
        let tran = {
            date: event.target.elements["date"].value,
            type: event.target.elements["type"].value,
            amount: event.target.elements["amount"].value,
            note: event.target.elements["note"].value
        };
        var arrayvar = this.state.transactions.slice();
        arrayvar.unshift(tran);

        this.setState({
            transactions: arrayvar,
            adding: false
        });
        console.log("ffff");
        console.log(this.state.transactions);
        event.target.reset();
    }

    changeType(typeName) {
        console.log(typeName);
        console.log(this.state.type);
        if (typeName != this.state.type) {
            console.log(this.state.type);
            this.setState({
                type: typeName
            });
        } else {
            return;
        }
    }

    changeAdding() {
        this.setState({
            adding: true
        });
        console.log("in");
    }

    render() {
        return (
            <div>
                <div className="App">
                    <header className="App-header">
                        <i
                            class="fa fa-bars fa-2x"
                            aria-hidden="true"
                            id="burger-bar"
                        />
                        <h1 className="App-title">Welcome to Midas</h1>
                    </header>
                    <p className="App-intro">
                        See your transactions or add a new transaction here!
                    </p>
                </div>
                <div className="types-btn">
                    <a
                        class="btn-floating btn-large waves-effect waves-light yellow"
                        onClick={this.changeType.bind(this, "all")}
                    >
                        <i class="fa fa-usd" aria-hidden="true" />
                    </a>
                    <a
                        class="btn-floating btn-large waves-effect waves-light yellow"
                        onClick={this.changeType.bind(this, "Food")}
                    >
                        <i class="fa fa-cutlery" aria-hidden="true" />
                    </a>
                    <a
                        class="btn-floating btn-large waves-effect waves-light yellow"
                        onClick={this.changeType.bind(this, "Travel")}
                    >
                        <i class="fa fa-plane" aria-hidden="true" />
                    </a>
                    <a
                        class="btn-floating btn-large waves-effect waves-light yellow"
                        onClick={this.changeType.bind(this, "Transportation")}
                    >
                        <i class="fa fa-bus" aria-hidden="true" />
                    </a>
                    <a
                        class="btn-floating btn-large waves-effect waves-light yellow"
                        onClick={this.changeType.bind(this, "Shopping")}
                    >
                        <i class="fa fa-shopping-bag" aria-hidden="true" />
                    </a>
                </div>
                {this.state.adding && (
                    <div class="money-form">
                        <form
                            class="col s12"
                            onSubmit={this.createTran.bind(this)}
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
                                <option value="Transportation">
                                    Transportation
                                </option>
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
                )}

                <div className="trans">
                    {this.state.transactions.map(d => {
                        if (
                            this.state.type == "all" ||
                            d.type == this.state.type
                        ) {
                            return <Transaction data={d} />;
                        }
                    })}
                </div>
                <a
                    class="btn-floating btn-large waves-effect waves-light yellow"
                    id="add-btn"
                    onClick={this.changeAdding.bind(this)}
                >
                    <i class="material-icons">+</i>
                </a>
            </div>
        );
    }
}

export default App;
