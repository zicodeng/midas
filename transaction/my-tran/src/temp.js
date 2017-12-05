                    <div className="money-form">
                        <form onSubmit={this.createTran.bind(this)}>
                            <p>Date</p>
                            <input
                                required
                                id="date"
                                placeholder="date"
                                type="text"
                                className="validate"
                            />
                            <p>Type</p>
                            <select required id="type">
                                <option value="Food">Food</option>
                                <option value="Transportation">
                                    Transportation
                                </option>
                                <option value="Shopping">Shopping</option>
                                <option value="Travel">Travel</option>
                            </select>
                            <p>Amount</p>
                            <input
                                required
                                id="amount"
                                placeholder="how much did you spend?"
                                type="text"
                                className="validate"
                            />
                            <p>Note</p>
                            <input
                                required
                                id="note"
                                placeholder="note"
                                className="validate"
                            />

                            <button type="submit" className="addTransaction">
                                Submit
                            </button>
                        </form>
                    </div>




                                        <p className="date">{this.props.data.date}</p>
                    <p className="note">{this.props.data.note}</p>
                    <p className="amount">- ${this.props.data.amount}</p>