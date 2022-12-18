import "./App.css";
import { BrowserRouter as Router } from "react-router-dom";

import Navigations from "@/routes";
function App() {
  return (
    <Router>
      <Navigations />
    </Router>
  );
}

export default App;
