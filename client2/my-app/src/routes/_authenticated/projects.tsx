import { Stack } from "@chakra-ui/react"
import { createFileRoute } from "@tanstack/react-router"
import { Container } from "@chakra-ui/react"
import ProjectForm from "@/components/project/ProjectForm"
import ProjectList from "@/components/project/ProjectList"

// const projectsQueryOptions = queryOptions({
//   queryKey: ["projects"],
//   queryFn: async () => {
// 			try {
// 				const res = await fetch(BASE_URL + "/projects", {
// 					// headers: { Authorization: `Bearer ${token}`},
// 					credentials: "include",
// 				});
// 				const data = await res.json();

// 				if (!res.ok) {
// 					throw new Error(data.error || "Something went wrong");
// 				}
// 				return data || [];
// 			} catch (error) {
// 				console.log(error);
// 			}
// 		},
// })

// const queryClient = useQueryClient()

export const Route = createFileRoute('/_authenticated/projects')({
  // loader: () => queryClient.ensureQueryData(projectsQueryOptions),
  component: Projects,
})

function Projects() {
    // const {
    //   data: { projects },
    // } = useSuspenseQuery(projectsQueryOptions)
    return (
    <Stack h="100vh">
      <Container>
        <ProjectForm />
        <ProjectList />
      </Container>
    </Stack>
  )
}
