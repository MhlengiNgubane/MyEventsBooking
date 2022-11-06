import *React from "react";
import ReactDOM from "react-dom";
import {HashRouter as Router, Route} from "react-router-dom";
import {EventListContainer} from "./components/EventListContainer";
import {Navigation} from "./components/Navigation";
import {EventBookingFormContainer} from "./components/EventBookingFormContainer";

class App extends React.Component<{}, {}> {
    render() {
        const eventList = () => <EventListContainer eventServiceURL="http://localhost:8181"/>;
        const eventBooking = ({match}:any) => <EventBookingFormContainer eventID={match.params.id}
            eventServiceURL="http://localhost:8181"
            bookingServiceURL="http://localhost:8282"/>;

        return <Router>
            <div>
                <Navigation brandName="MyEvents"/>
                <div className="container">
                    <h1>My Events</h1>

                    <Route exact path="/" component={eventList}/>
                    <Route path="/events/:id/book" component={eventBooking}/>
                </div>
            </div>
        </Router>
    }
}

ReactDOM.render(
    <App/>,
    document.getElementById("myevents-app")
);