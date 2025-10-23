import { BASE_URL } from "@/main";
import { Input, Button, Spinner, Stack, Text } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";

const LoginForm = () => {
    const [login, setLogin] = useState("neo");
    const [password, setPassword] = useState("keanu");

    const navigate = useNavigate();

    const queryClient = useQueryClient();

    const { mutate: auth, isPending: isCreating } = useMutation({
        mutationKey: ["auth"],
        mutationFn: async (e: React.FormEvent) => {
            e.preventDefault();
            try {
                const res = await fetch(BASE_URL + `/auth`, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({ user: login, pwd: password }),
                });
                const data = await res.json();

                if (!res.ok) {
                    throw new Error(data.error || "Something went wrong");
                }

                setLogin("");
                setPassword("");
                navigate({
                    to: '/',
                })
                return data;
            } catch (error: any) {
                throw new Error(error);
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ["auth"] });
        },
        onError: (error: any) => {
            alert(error.message);
        },
    });

    return (
        <form onSubmit={auth}>
            <Stack gap={2}>
                <Text>
                    Login
                </Text>
                <Input
                    type='text'
                    value={login}
                    onChange={(e) => setLogin(e.target.value)}
                    // ref={(input) => input && input.focus()}
                />
                <Text>
                    Password
                </Text>
                <Input
                    type='text'
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    // ref={(input) => input && input.focus()}
                />
                <Button
                    mx={2}
                    type='submit'
                    _active={{
                        transform: "scale(.97)",
                    }}
                >
                    {isCreating ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
                </Button>
            </Stack>
        </form>
    );
};
export default LoginForm;