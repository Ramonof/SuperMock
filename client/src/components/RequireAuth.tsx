import { useLocation, Outlet, Navigate } from "react-router";
import useAuth from "../hooks/useAuth";

const RequireAuth = ({ allowedRoles }: any) => {
    const { auth } = useAuth();
    const location = useLocation();
    console.log(auth)
    console.log(location)
    console.log(auth.accessToken)

    return (
        auth?.roles?.find((role: any) => allowedRoles?.includes(role))
            ? <Outlet />
            : auth?.user
                ? <Navigate to="/unauthorized" state={{ from: location }} replace />
                : <Navigate to="/login" state={{ from: location }} replace />
    );
}

export default RequireAuth;