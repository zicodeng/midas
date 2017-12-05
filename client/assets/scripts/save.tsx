import * as React from 'react';
let features = [
    {
        featureName: 'Transaction',
        featureDescription:
            'You can easily input income and expenses and check it back when you need it',
        featureImage:
            'http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg',
        featurePageLink: ''
    },

    {
        featureName: 'Scan Receipt',
        featureDescription: 'Digitize texts using your camera',
        featureImage:
            'http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg',
        featurePageLink: ''
    },

    {
        featureName: 'Graph Report',
        featureDescription: "See graphs to get a better idea of where you're spending youre money",
        featureImage:
            'http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg',
        featurePageLink: ''
    },

    {
        featureName: 'Set Budget',
        featureDescription: 'Set a budget by customizing your settings and Midas will do the rest',
        featureImage:
            'http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg',
        featurePageLink: ''
    },

    {
        featureName: 'Saving Goal',
        featureDescription: 'Create a saving goal and Midas will help you achieve it',
        featureImage:
            'http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg',
        featurePageLink: ''
    }
];

class SaveBudget extends React.Component<any, any> {
    constructor(props) {
        super(props);
        this.state = {
            budget: 180,
            spent: 50,
            saving_name:"",
            saving_goal: 0,
            income: 0,
            saving_budget: 0,
            edit:false
        };
    
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
      }
    
      handleChange(event) {
        var name = event.target.name;
        var value = event.target.value;
        this.state[name] = value;
        // this.setState({value: event.target.value});
    this.forceUpdate();   
      }
      click(e){
          e.preventDefault;
      }
    validate(){
        return (this.state.saving_name.length >0) && (this.state.saving_goal >0) &&
        (this.state.income >0) && (this.state.saving_budget > 0);

    }
      handleSubmit(event) {
        alert('A name was submitted: ' + this.state.value);
        event.preventDefault();
      }
    componentDidMount() {
        this.setState({ name: this.props.name });
    }
    public render() {
        return (
            <div className="content" >
                <div className="header">
                    <a>
                        <img src="http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg" />
                    </a>
                    <button className="sign_out_button">Sign Out</button>
                </div>
                <div className="greeting budget">
                    <h1>Budget</h1>
                    <hr />
                    <h2>Budget for this month...</h2>
                    {(this.state.budget > 0) &&
                    <div>
                    <progress max={this.state.budget} value={this.state.spent}>{this.state.spent}</progress>
                    <div className="budget_text" >You've spent $ {this.state.spent} out of $ {this.state.budget} </div>
                    { (this.state.spent > this.state.budget) &&
                        <div className="spending-alert">
                            <p >You've exceeded your spending goal!</p>
                        </div>
                    }

                    </div>
                    }
                    {(this.state.budget <= 0) &&
                        <div className="spending-alert">
                            <p >Set your spending goal!</p>
                        </div>
                    }
                    <hr />
                </div>
                <div className="saving budget">
                    <h2>Budget for December...</h2>
                    {this.state.edit &&
                    <form >
                        <div className="form-group">
                            <label className="control-label budget">What's your monthly budget?</label>
                            <input type="text" value={this.state.budget} name="budget" placeholder="Budget" id="monthly_budget" className="form-control" onChange={this.handleChange} />
                        </div>
                        <button disabled={this.state.budget <0}>Save</button>
                        <div className="form-group">
                            <label className="control-label budget">Anything you want to save for? What's it called?</label>
                            <input type="text" value={this.state.saving_name} name="saving_name" id="saving_name" className="form-control" onChange={this.handleChange}/>
                            {(this.state.saving_name.length <= 0) &&
                             <p className="input-alert">Must enter name to start saving goal</p>
                            }
                        </div>  
                        { (this.state.saving_name.length >0) &&
                         <div className="form-group">
                            <label className="control-label budget">How much does it cost?</label>
                            <input type="text" value={this.state.saving_goal} name="saving_goal" placeholder="Goal" id="saving_goal" className="form-control" onChange={this.handleChange} />
                            {(this.state.saving_goal <= 0) &&
                             <p className="input-alert">Must enter the goal's cost</p>
                            }
                        </div> 
                        }
                        {(this.state.saving_goal >0) &&
                        <div className="form-group">
                            <label className="control-label budget">How much is your monthly income</label>
                            <input type="text" value={this.state.income} name="income" placeholder="Income" id="income" className="form-control" onChange={this.handleChange}/>
                            {(this.state.income <= 0) &&
                             <p className="input-alert">Must enter your income</p>
                            }
                        </div> 
                        }
                        {(this.state.income >0) &&
                        
                        <div className="form-group">
                            <label className="control-label budget">How much do you want to save per month?</label>
                            <input type="text" value={this.state.saving_budget} name="saving_budget" placeholder="Saving Budget" id="saving_budget" className="form-control" onChange={this.handleChange}/>
                            {(this.state.saving_budget <= 0) &&
                             <p className="input-alert">Must enter your monthly saving budget</p>
                            }
                        </div>
                        }  
                        <button disabled={!this.validate} onClick={this.handleSubmit}>Save</button>
         
                    </form>

                    }
                    {!this.state.edit &&
                    <div>
                        <h2>$ {this.state.budget}</h2>
                        <h2>Special spending budget</h2>
                        <h2>$ {this.state.saving_goal} {this.state.saving_name}</h2>
                        <h2>So far you've saved $ {this.state.saving_budget}</h2>

                    </div>

                    }
                </div>
            </div>
       
        );
    }
}

// //set up buget
// //if spending if above budget then notification pops up
// //saving goal 
// //saving increment
// //income
// //show time til saved


export default SaveBudget;