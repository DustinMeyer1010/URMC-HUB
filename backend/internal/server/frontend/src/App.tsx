import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Home from "./pages/Home";
import Search from "./pages/Search";
import { ProtectedRoute } from "./components/ProtectedRoute";
import { useAuth } from "./components/Authentication";
import Login from "./pages/Login";
import Nav from "./components/Nav";

function App() {

  const { loading, isLoggedIn } = useAuth();

  if (loading || isLoggedIn === null) {
    return <div>Loading...</div>;
  }

  return (
      <Router>
        <Nav/>

        <Routes>
          <Route path="/" element={ <Home /> } />
            <Route path="/search" element={
              <ProtectedRoute>
                <Search />
              </ProtectedRoute>
              } 
            />
          <Route path="/login" element={ <Login/> }/>
        </Routes>
      </Router>
    );
  }

  export default App;