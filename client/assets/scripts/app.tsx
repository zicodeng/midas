import * as React from 'react';

import SavingBudget from './save';

class App extends React.Component<any, any> {
    private count: number = 0;

    constructor(props, context) {
        super(props, context);
        this.state = {
            logged: true,
            save: false,
            name: 'temp',
            data: ['']
        };
    }

    public render() {
        let landingPage = <HomePage />;
        switch (this.state.logged) {
            case true:
                landingPage = <WelcomePage name={this.state.name} />;
        }
        if (this.state.save) {
            landingPage = <SavingBudget name={this.state.name} />;
        }
        return (
            <div>
                <button onClick={e => this.handleClickBtn()}>Save</button>
                {landingPage}
                <MainFooter />
            </div>
        );
    }

    private handleClickBtn = () => {
        let save = this.state.save;
        this.setState({
            save: !save
        });
    };
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

class HomePage extends React.Component<any, any> {
    public render() {
        return (
            <div className="content">
                <div className="header">
                    <a className="logo">
                        <img src="http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg" />
                        <h1>Midas</h1>
                    </a>
                    <a className="log_in_button">Join</a>
                </div>
                <div className="greeting">
                    <h1>Midas</h1>
                    <div className="greeting_text">
                        Helping you to manage your money all in one place
                    </div>
                </div>
                <div className="features">
                    {features.map((item, index) => (
                        <a
                            className="h_feature_item"
                            id={(() => {
                                if (index % 2 == 1) {
                                    return 'odd';
                                }
                            })()}
                            href={item.featurePageLink}
                            key={index}
                        >
                            <img className="feature_img" src={item.featureImage} />
                            <div className="feature_description">
                                <h3 className="h_feature_header">{item.featureName}</h3>
                                <div className="feature_text">{item.featureDescription}</div>
                            </div>
                        </a>
                    ))}
                </div>
                <div className="button_wrapper">
                    <a className="bottom_log_in_button"> Join </a>
                </div>
            </div>
        );
    }
}
class WelcomePage extends React.Component<any, any> {
    constructor() {
        super();
        this.state = {
            name: ''
        };
    }

    componentDidMount() {
        this.setState({ name: this.props.name });
    }

    public render() {
        return (
            <div className="content">
                <div className="header">
                    <a className="logo">
                        <img src="http://students.washington.edu/kpham97/websitetest/static/media/logo.5d5d9eef.svg" />
                    </a>
                    <a className="sign_out_button">Sign Out</a>
                </div>
                <div className="greeting">
                    <h1>Hello</h1>
                    <div className="greeting_text">Welcome back, {this.state.name}</div>
                </div>
                <div className="features">
                    {features.map((item, index) => (
                        <a className="feature_item" href={item.featurePageLink} key={index}>
                            <h3 className="feature_header">{item.featureName}</h3>
                            <img
                                className="arrow"
                                src="http://students.washington.edu/kpham97/arrow-to-right.jpg"
                            />
                        </a>
                    ))}
                </div>
                <div className="button_wrapper">
                    <button className="bottom_sign_out_button"> Sign Out </button>
                </div>
            </div>
        );
    }
}

class MainFooter extends React.Component<any, any> {
    public render() {
        return (
            <div className="footer">
                <a>Home&emsp;|&emsp;</a>
                <a>Features&emsp;|&emsp;</a>
                <a>Company&emsp;|&emsp;</a>
                <a>Contact</a>
            </div>
        );
    }
}
export default App;
