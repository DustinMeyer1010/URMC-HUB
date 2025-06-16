import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Home from "./pages/Home";
import Search from "./pages/Search";

function App() {
  return (
    <Router>
      <nav style={{ display: "flex", gap: "1rem", marginBottom: "1rem", position: 'fixed', top: "10px", left: "10px", zIndex: "10"}}>
        <Link to="/">Home</Link>
        <Link to="/search">Search</Link>
        <Link to="/SummaryIndex">System Summary Index</Link>
      </nav>

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/search" element={<Search />} />
      </Routes>
    </Router>
  );
}

export default App;