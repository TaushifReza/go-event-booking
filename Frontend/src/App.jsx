import { BrowserRouter, Outlet, Route, Routes } from "react-router-dom";

import HomePage from "./pages/HomePage";
import Login from "./pages/Login";
import Register from "./pages/Register";
import { AuthProvider } from "./context/AuthContext";
import Dashboard from "./pages/Dashboard";
import ProtectedRoute from "./pages/ProtectedRoute";
import Layout from "./pages/Layout";
import NotFound from "./pages/NotFound";

function App() {
    return (
        <>
            <AuthProvider>
                <BrowserRouter>
                    <Routes>
                        {/* Public Routes */}
                        <Route index element={<HomePage />} />
                        <Route path="login" element={<Login />} />
                        <Route path="register" element={<Register />} />

                        {/* Protected Route Group */}
                        <Route
                            element={
                                <ProtectedRoute>
                                    <Layout />
                                </ProtectedRoute>
                            }
                        >
                            <Route path="dashboard" element={<Dashboard />} />
                        </Route>

                        <Route path="*" element={<NotFound />} />
                    </Routes>
                </BrowserRouter>
            </AuthProvider>
        </>
    );
}

export default App;
