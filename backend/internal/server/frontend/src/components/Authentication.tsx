import axios from "axios";
import { createContext, useContext, useEffect, useState } from "react";
import type { ReactNode } from "react";
import devConfig from "../../devconfig.json"

type Auth = {
    isLoggedIn: boolean | null;
    loading: boolean;
    login: () => void;
    logout: () => void;
}

const AuthContext = createContext<Auth | undefined>(undefined)

export function AuthProvider({ children }: { children: ReactNode }) {
    const [isLoggedIn, setLoggedIn] = useState<boolean | null>(null);
    const [loading, setloading] = useState(true);

    const login = () => setLoggedIn(true);
    const logout = () => setLoggedIn(false);

        useEffect(() => {
            const checkLoginStatus = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/verify`, {});
                if (response.status == 200) {
                    login()
                }    
                else {
                    logout()
                }
            } catch (error) {
                    setTimeout(() => {
                        logout();
                    }, 200);
                    return;
            }finally {
                setloading(false)
            };
        }

        const requireLogin = devConfig.login.require === "true"
        
        if (requireLogin){
            checkLoginStatus();
<<<<<<< Updated upstream
            /*
            login()
            setloading(false)
            */
=======
        }
        else {
            login()
            setloading(false)
         
        }
           
>>>>>>> Stashed changes
            }, []);

    return (
        <AuthContext.Provider value={{ isLoggedIn, loading, login, logout }}>
            {children}
        </AuthContext.Provider>
    )
}

export function useAuth() {
    const context = useContext(AuthContext)
    if (!context) throw new Error("useAuth must be used within an AuthProvider")
    return context;
}