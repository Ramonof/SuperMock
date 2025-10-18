import ProjectName from "@/components/generic/ProjectName"
import RestStubList from "@/components/stub/rest/RestStubList"
import StubsInfo from "@/components/stub/rest/StubsInfo"
import { Stack, Container } from "@chakra-ui/react"
import { createFileRoute } from "@tanstack/react-router"

export const Route = createFileRoute('/_authenticated/project_/$projectId/rest/stubs_/$stubId')({
  component: RestStubs,
})

function RestStubs() {
  const { projectId } = Route.useParams()
  return (
      <>
        <ProjectName ProjectId={projectId}/>
        <StubsInfo ProjectId={projectId}/>
        <RestStubList ProjectId={projectId}/>
      </>
  )
}