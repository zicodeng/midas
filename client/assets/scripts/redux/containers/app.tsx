import * as React from "react";
import { connect } from "react-redux";

import * as UserActions from "scripts/redux/actions/user-actions";
import {Component} from "react";

@connect((Store) => {
    return {
        user: Store.user
    };
})

class App extends React.Component<any, any> {
    private count: number = 0;

    constructor(props, context) {
        super(props, context);
        this.state = {
            logged: false,
            name: "temp",
            data: [""],
        }
    }

    public render() {
        let landingPage = <HomePage />
        switch(this.state.logged){
            case true:
                return(landingPage = <WelcomePage name={this.state.name} />);
        }
        return (
            <div>
                {landingPage}
                <MainFooter />
            </div>
        )
    }

    private handleBtnClicked(e) {
        this.props.dispatch(UserActions.setAge(this.count));
        this.count += 2;
    }
}

/*
function LandingPage(props){
    const logged = props.logged;
    if (logged){
        return <WelcomePage username={props.username} />;
    }

    return <HomePage/>;
}*/

let features = [

    {featureName : "Transaction", featureDescription: "You can easily input income and expenses and check it back when you need it",
        featureImage : "http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg", featurePageLink:""},

    {featureName : "Scan Receipt", featureDescription: "Digitize texts using your camera",
        featureImage : "http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg", featurePageLink:""},

    {featureName : "Graph Report", featureDescription: "See graphs to get a better idea of where you're spending youre money",
        featureImage : "http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg", featurePageLink:""},

    {featureName : "Set Budget", featureDescription: "Set a budget by customizing your settings and Midas will do the rest",
        featureImage : "http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg", featurePageLink:""},

    {featureName : "Saving Goal", featureDescription: "Create a saving goal and Midas will help you achieve it",
        featureImage : "http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg", featurePageLink:""},
];

class HomePage extends Component{
    public render() {
        return (
            <div className="content">
                <div className="header">
                    <a>
                        <img src = "http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg"/>
                        <h1>Midas</h1>
                    </a>
                    <button className="log_in_button">Join</button>
                </div>
                <div className="greeting">
                    <h1>Midas</h1>
                    <div className="greeting_text">
                        Helping you to manage money all in one place
                    </div>
                </div>
                <div className="h_features">
                    {features.map((item, index) => (
                        <a className="h_feature_item" href={item.featurePageLink} key = {index}>
                            <img className="h_feature_img" src = {item.featureImage}/>
                            <h3 className="h_feature_header">{item.featureName}</h3>
                            <div className="h_feature_text">{item.featureDescription}</div>
                        </a>
                        ))
                    }
                </div>
                <button className="bottom_log_in_button"> Join </button>
            </div>
        )
    }

}
class WelcomePage extends Component{

    constructor(){
        super();
        this.state = {
            name = ""
        };
    }

    componentDidMount{
        setState(name: this.props.name);
    }

    public render() {
        return (
            <div className="content">
                <div className="header">
                    <a>
                        <img src ="http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg"/>
                    </a>
                    <button className="sign_out_button">Sign Out</button>
                </div>
                <div className="greeting">
                    <h1>Hello</h1>
                    <div className="greeting_text">
                       Welcome back, {this.props.name}
                    </div>
                </div>
                <div className="w_features">
                    {features.map((item, index) => (
                        <a className="w_feature_item" href={item.featurePageLink} key = {index}>
                            <h3 className="w_feature_header">{item.featureName}</h3>
                            <img className="arrow" src = "http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg"/>
                        </a>
                    ))
                    }
                </div>
                <button className="bottom_log_in_button"> Join </button>
            </div>
        )
    }

}

class MainFooter extends Component {
    public render(){
        return(
            <div>
                <a> Home |</a>
                <a> Features |</a>
                <a> Company |</a>
                <a> Contact |</a>
            </div>
        )
    }
}
export default App;


