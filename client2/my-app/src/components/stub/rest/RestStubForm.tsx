import { BASE_URL } from "@/main";
import { Flex, Input, Button, Spinner, Stack, Text } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";

const RestStubForm = ({ ProjectId }: { ProjectId: string }) => {
	const [newRestStubName, setNewRestStubName] = useState("");
	const [newRestStubPath, setNewRestStubPath] = useState("");
	const [newRestStubResponseBody, setNewRestStubResponseBody] = useState("");

	const navigate = useNavigate();

	const queryClient = useQueryClient();

	const { mutate: createRestStub, isPending: isCreating } = useMutation({
		mutationKey: ["createRestStub"],
		mutationFn: async (e: React.FormEvent) => {
			e.preventDefault();
			try {
				const res = await fetch(BASE_URL + `/projects/` + ProjectId + `/stub`, {
					method: "POST",
					headers: {
						"Content-Type": "application/json",
					},
					body: JSON.stringify({ name: newRestStubName, path: newRestStubPath, response_body: newRestStubResponseBody }),
				});
				const data = await res.json();

				if (!res.ok) {
					throw new Error(data.error || "Something went wrong");
				}

				setNewRestStubName("");
				setNewRestStubPath("");
				setNewRestStubResponseBody("");
				navigate({
					to: '/project/$projectId/rest/stubs',
					params: { projectId: ProjectId },
				})
				return data;
			} catch (error: any) {
				throw new Error(error);
			}
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["getRestStubs"] });
		},
		onError: (error: any) => {
			alert(error.message);
		},
	});

	return (
		<form onSubmit={createRestStub}>
			<Stack gap={2}>
				<Text>
					Name
				</Text>
				<Input
					type='text'
					value={newRestStubName}
					onChange={(e) => setNewRestStubName(e.target.value)}
					// ref={(input) => input && input.focus()}
				/>
				<Text>
					Path
				</Text>
				<Input
					type='text'
					value={newRestStubPath}
					onChange={(e) => setNewRestStubPath(e.target.value)}
					// ref={(input) => input && input.focus()}
				/>
				<Text>
					Body
				</Text>
				<Input
					type='text'
					value={newRestStubResponseBody}
					onChange={(e) => setNewRestStubResponseBody(e.target.value)}
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
export default RestStubForm;