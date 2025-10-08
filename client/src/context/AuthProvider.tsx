import { createContext, useState, type Dispatch, type ReactNode, type SetStateAction } from "react";

interface AuthState {
    user: string
    pwd: string
    roles: any
    accessToken: any
}

interface AuthContextType {
    auth: AuthState;
    setAuth: Dispatch<SetStateAction<AuthState>>;
}

const AuthContext = createContext<AuthContextType>({
    auth: { user: "", pwd: "", roles: null, accessToken: null },
    setAuth: () => {},
});

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [auth, setAuth] = useState<AuthState>({
        user: "", pwd: "", roles: null, accessToken: null
    });
    
    return (
        <AuthContext.Provider value={{ auth, setAuth }}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthContext;