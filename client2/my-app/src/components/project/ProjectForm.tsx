import { BASE_URL } from "@/main";
import apiClient from "@/utils/request";
import { Flex, Input, Button, Spinner } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";

const ProjectForm = () => {
	const [newProject, setNewProject] = useState("");

	const queryClient = useQueryClient();

	const { mutate: createProject, isPending: isCreating } = useMutation({
		mutationKey: ["createProject"],
		mutationFn: async (e: React.FormEvent) => {
			e.preventDefault();
			try {
				const res = await apiClient({
					url: BASE_URL + `/projects`, 
					method: 'POST',
					data: JSON.stringify({ name: newProject }),
				});
				setNewProject("");
				return res;
			} catch (error: any) {
				throw new Error(error);
			}
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["projects"] });
		},
		onError: (error: any) => {
			alert(error.message);
		},
	});

	return (
		<form onSubmit={createProject}>
			<Flex gap={2}>
				<Input
					type='text'
					value={newProject}
					onChange={(e) => setNewProject(e.target.value)}
					ref={(input) => input && input.focus()}
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
			</Flex>
		</form>
	);
};
export default ProjectForm;