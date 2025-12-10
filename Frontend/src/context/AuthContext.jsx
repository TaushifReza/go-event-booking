import { createContext, useContext, useState } from "react";

import { loginUser } from "../api/auth";
import { fetchUserProfile } from "../api/user";

// 1.create the context object
const AuthContext = createContext();

// 2. create Provider Component
function AuthProvider({ children }) {
    const [user, setUser] = useState(null);
    const [accessToken, setAccessToken] = useState("");
    const [refreshToken, setRefreshToken] = useState("");
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [isLoading, setIsLoading] = useState(false);

    const login = async (email, password) => {
        setIsLoading(true);
        const loginResult = await loginUser(email, password);

        if (loginResult.success) {
            const newAccessToken = loginResult.data.access_token;
            const newRefreshToken = loginResult.data.refresh_token;
            setAccessToken(newAccessToken);
            setRefreshToken(newRefreshToken);

            // fetch user profile
            const profileResult = await fetchUserProfile(newAccessToken);

            console.log("Profile res: ", profileResult);

            if (profileResult.success) {
                setUser(profileResult.data);
                setIsLoading(false);
                setIsAuthenticated(true);
                return { success: true };
            } else {
                setIsLoading(false);
                logout();
                return {
                    success: false,
                    message: profileResult.message,
                    error: profileResult.error,
                };
            }
        }
        setIsLoading(false);
        return {
            success: false,
            message: loginResult.message,
            error: loginResult.error,
        };
    };

    const logout = () => {
        setUser(null);
        setAccessToken("");
        setRefreshToken("");
    };

    const value = {
        user,
        accessToken,
        refreshToken,
        isAuthenticated,
        isLoading,
        login,
        logout,
    };

    return (
        <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
    );
}

// 3. Custom hook to easily consume the context
function useAuth() {
    const context = useContext(AuthContext);
    if (context === undefined)
        throw new Error("AuthContext was used outside AuthProvider");

    return context;
}

export { AuthProvider, useAuth };
