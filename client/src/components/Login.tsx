import { Button, Flex, Input, Link, Spinner } from "@chakra-ui/react"
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { useEffect, useRef, useState } from "react";
import { IoMdAdd } from "react-icons/io";
import { BASE_URL } from "../App";
import { useNavigate, useLocation } from "react-router";
import useAuth from "../hooks/useAuth";

export type Login = {
	id: number;
	name: string;
	created_at: string;
};

// const Login = () => {
//     const [newProject, setNewProject] = useState("");

// 	const queryClient = useQueryClient();

//     // const { mutate: login, isPending: isLogedIn } = useQuery({
//     //     mutationKey: ["login"],
// 	// 	mutationFn: async (e: React.FormEvent) => {
// 	// 		e.preventDefault();
            
//     //     }
//     // });

//     const { data: login, isLoading } = useQuery<Login[]>({
//             queryKey: ["projects"],
//             queryFn: async () => {
//                 try {
//                     const res = await fetch(BASE_URL + `/auth`, {
//                         method: "POST",
//                         headers: {
//                             "Content-Type": "application/json",
//                         },
//                         // credentials: "include",
//                         body: JSON.stringify({ name: newProject }),
//                     });
//                     const data = await res.json();

//                     if (!res.ok) {
//                         throw new Error(data.error || "Something went wrong");
//                     }
//                     return data || [];
//                 } catch (error) {
//                     console.log(error);
//                 }
//             },
//         });

//     return (
        
//         <form onSubmit={auth}>
//                     <Flex gap={2}>
//                         <Input
//                             type='text'
//                             value={newProject}
//                             onChange={(e) => setNewProject(e.target.value)}
//                             ref={(input) => input && input.focus()}
//                         />
//                         <Button
//                             mx={2}
//                             type='submit'
//                             _active={{
//                                 transform: "scale(.97)",
//                             }}
//                         >
//                             {isCreating ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
//                         </Button>
//                     </Flex>
//                 </form>

//     )
// }

const Login = () => {
    const { auth, setAuth } = useAuth();

    const navigate = useNavigate();
    const location = useLocation();
    const from = location.state?.from?.pathname || "/";

    const userRef = useRef<HTMLInputElement>(null);
    const errRef = useRef<HTMLInputElement>(null);

    const [user, setUser] = useState('');
    const [pwd, setPwd] = useState('');
    const [errMsg, setErrMsg] = useState('');

    useEffect(() => {
        if (userRef.current) {
            userRef.current.focus();
        }
    }, [])

    useEffect(() => {
        setErrMsg('');
    }, [user, pwd])

    const handleSubmit = async (e: { preventDefault: () => void; }) => {
        e.preventDefault();

        try {
            const response = await axios.post(BASE_URL + `/auth`,
                JSON.stringify({ user, pwd }),
                {
                    headers: { 'Content-Type': 'application/json' },
                    withCredentials: true
                }
            );
            console.log(JSON.stringify(response?.data));
            //console.log(JSON.stringify(response));
            const accessToken = response?.data?.accessToken;
            const roles = response?.data?.roles;
            //TODO
            const userData = { user: user, pwd: pwd, roles: roles, accessToken: accessToken };
            setAuth(userData);
            console.log("auth")
            console.log(setAuth)
            console.log(auth)
            setUser('');
            setPwd('');
            navigate(from, { replace: true });
        } catch (err : any) {
            if (!err?.response) {
                setErrMsg('No Server Response');
            } else if (err.response?.status === 400) {
                setErrMsg('Missing Username or Password');
            } else if (err.response?.status === 401) {
                setErrMsg('Unauthorized');
            } else {
                setErrMsg('Login Failed');
            }
            if (errRef.current) {
                errRef.current.focus();
            }
        }
    }

    return (

        <section>
            <p ref={errRef} className={errMsg ? "errmsg" : "offscreen"} aria-live="assertive">{errMsg}</p>
            <h1>Sign In</h1>
            <form onSubmit={handleSubmit}>
                <label htmlFor="username">Username:</label>
                <input
                    type="text"
                    id="username"
                    ref={userRef}
                    autoComplete="off"
                    onChange={(e) => setUser(e.target.value)}
                    value={user}
                    required
                />

                <label htmlFor="password">Password:</label>
                <input
                    type="password"
                    id="password"
                    onChange={(e) => setPwd(e.target.value)}
                    value={pwd}
                    required
                />
                <button>Sign In</button>
            </form>
            <p>
                Need an Account?<br />
                <span className="line">
                    <Link to="/register">Sign Up</Link>
                </span>
            </p>
        </section>

    )
}

export default Login