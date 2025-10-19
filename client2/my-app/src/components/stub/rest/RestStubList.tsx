import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import { useQuery } from "@tanstack/react-query";
import { BASE_URL } from "@/main";
import RestStubItem from "./RestStubItem";
import apiClient from "@/utils/request";


export type RestStub = {
	id: number;
	name: string;
    project_id: string;
	created_at: string;
    path: string;
	method: string;
    response_body: string;
};

const RestStubList = ({ ProjectId }: { ProjectId: string }) => {
	const { data: projects, isLoading } = useQuery<RestStub[]>({
		queryKey: ["projects"],
		queryFn: async () => {
			try {
				const res = await apiClient({
					url: BASE_URL + `/projects/` + ProjectId + `/stub`, 
					method: 'get'
				});
				const data = res.data

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
				 
			</Text>
			{isLoading && (
				<Flex justifyContent={"center"} my={4}>
					<Spinner size={"xl"} />
				</Flex>
			)}
			<Stack gap={3}>
				{projects?.map((Project) => (
					<RestStubItem key={Project.id} Project={Project} />
				))}
			</Stack>
		</>
	);
};
export default RestStubList;