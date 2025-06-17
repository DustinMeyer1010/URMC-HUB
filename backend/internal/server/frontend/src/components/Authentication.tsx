import { createContext, useContext, useState } from "react";
import type { ReactNode } from "react";

type Auth = {
    isLoggedIn: boolean;
    login: () => void;
    logout: () => void;
}

const AuthContext = createContext<Auth | undefined>(undefined)

export function AuthProvider({ children }: { children: ReactNode }) {
    const [isLoggedIn, setLoggedIn] = useState(false);

    const login = () => setLoggedIn(true);
    const logout = () => setLoggedIn(false);

    return (
        <AuthContext.Provider value={{ isLoggedIn, login, logout }}>
            {children}
        </AuthContext.Provider>
    )
}

export function useAuth() {
    const context = useContext(AuthContext)
    if (!context) throw new Error("useAuth must be used within an AuthProvider")
    return context;
}