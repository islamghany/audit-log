import "./App.css";
import { BrowserRouter as Router, Link } from "react-router-dom";

import Navigations from "@/routes";
function App() {
  return (
    <Router>
      <nav className="my-8 space-x-4">
        <Link to="/">Dashboard</Link>
        <Link to="/login">Login</Link>
        <Link to="/register">Register</Link>
      </nav>

      <Navigations />
    </Router>
  );
}

export default App;
