// src/pages/ProtectedRoute.jsx (Improved)

import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";
import { useEffect } from "react";

function ProtectedRoute({ children }) {
    const { isAuthenticated, isLoading } = useAuth();
    const navigate = useNavigate();

    useEffect(
        function () {
            if (!isLoading && !isAuthenticated) {
                navigate("/login", { replace: true });
            }
        },
        [isAuthenticated, isLoading, navigate]
    );

    if (isLoading) {
        return (
            <div className="flex justify-center items-center h-screen">
                <p>Loading user session...</p>
            </div>
        );
    }

    return isAuthenticated ? children : null;
}

export default ProtectedRoute;
