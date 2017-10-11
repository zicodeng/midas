// Can't use import React from "react",
// because it is not exported as default.
import * as React from "react";
import * as ReactDOM from "react-dom";

import { Provider } from "react-redux";
import Store from "scripts/redux/store";

// Import container.
import App from "scripts/redux/containers/app";

// Import bootstrap.
// import "bootstrap/dist/css/bootstrap.min.css";

// Import custom stylesheets.
import "stylesheets/entries/app.scss";

ReactDOM.render(
    <Provider store={ Store }>
        <App />
    </Provider>,
    document.getElementById("app")
);