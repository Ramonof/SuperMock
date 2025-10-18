import { BASE_URL } from "@/main";
import apiClient from "@/utils/request";
import { Container, Text } from "@chakra-ui/react";
import type { RestStub } from "../stub/rest/RestStubList";
import type { Project } from "../project/ProjectList";
import { useQuery } from "@tanstack/react-query";

export default function ProjectName({ ProjectId }: { ProjectId: string }) {
    const { data: project, isLoading } = useQuery<Project>({
		queryKey: ["projectName"],
		queryFn: async () => {
			try {
				const response = await apiClient({
					url: BASE_URL + `/projects/` + ProjectId, 
					method: 'get'
				});
				// const res = response.data
				const data = response.data

				// if (!res.ok) {
				// 	throw new Error(data.error || "Something went wrong");
				// }
				return data || null;
			} catch (error) {
				console.log(error);
			}
		},
	});


    return (
        <Container maxW={"900px"}>
            <Text
                fontSize={"4xl"}
                textTransform={"uppercase"}
                fontWeight={"bold"}
                textAlign={"center"}
                my={2}
                bgGradient='linear(to-l, #0bf827ff, #4000ffff)'
                bgClip='text'
            >
                {project?.name}
            </Text>
        </Container>
    )
}