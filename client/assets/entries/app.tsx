// Can't use import React from "react",
// because it is not exported as default.
import * as React from 'react';
import * as ReactDOM from 'react-dom';

// Import container.
import App from 'scripts/app';

// Import bootstrap.
// import "bootstrap/dist/css/bootstrap.min.css";

// Import custom stylesheets.
import 'stylesheets/entries/app.scss';

ReactDOM.render(<App />, document.getElementById('app'));
