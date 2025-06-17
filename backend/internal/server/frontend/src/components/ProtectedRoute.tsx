import { Navigate } from "react-router-dom";
import { useAuth } from "./Authentication";
import type { ReactNode } from "react";


export function ProtectedRoute({ children }: { children: ReactNode }) {
    const { isLoggedIn } = useAuth()


    if (!isLoggedIn) {
        return <Navigate to="/login" replace />;
    }

    return <>{ children }</>
}