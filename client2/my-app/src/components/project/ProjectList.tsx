import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import { useQuery } from "@tanstack/react-query";
import ProjectItem from "./ProjectItem";
import { BASE_URL } from "@/main";
import axios from "@/utils/request";
import apiClient from "@/utils/request";

export type Project = {
	id: number;
	name: string;
	created_at: string;
};

const ProjectList = () => {
	const { data: projects, isLoading } = useQuery<Project[]>({
		queryKey: ["projects"],
		queryFn: async () => {
			try {
				const response = await apiClient({
					url: BASE_URL + "/projects", 
					method: 'get'
				});
				// const res = response.data
				const data = response.data

				// if (!res.ok) {
				// 	throw new Error(data.error || "Something went wrong");
				// }
				return data || [];
			} catch (error) {
				console.log(error);
			}
		},
	});

	return (
		<>
			<Text
				fontSize={"4xl"}
				textTransform={"uppercase"}
				fontWeight={"bold"}
				textAlign={"center"}
				my={2}
				bgGradient='linear(to-l, #0bf827ff, #4000ffff)'
				bgClip='text'
			>
				Projects
			</Text>
			{isLoading && (
				<Flex justifyContent={"center"} my={4}>
					<Spinner size={"xl"} />
				</Flex>
			)}
			{!isLoading && projects?.length === 0 && (
				<Stack alignItems={"center"} gap='3'>
					<Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
						All tasks completed! 🤞
					</Text>
					<img src='/go.png' alt='Go logo' width={70} height={70} />
				</Stack>
			)}
			<Stack gap={3}>
				{projects?.map((Project) => (
					<ProjectItem key={Project._id} Project={Project} />
				))}
			</Stack>
		</>
	);
};
export default ProjectList;