import React, { Component } from "react";
import "./Transaction.css";

class Transaction extends Component {
    render() {
        return (
            <div className="allTransaction">
                <div className="tran-sec">
                    <span className="itemize">{this.props.data.date}</span>
                    <span className="itemize">{this.props.data.note}</span>
                    <span className="text">-${this.props.data.amount}</span>
                </div>
            </div>
        );
    }
}

export default Transaction;
