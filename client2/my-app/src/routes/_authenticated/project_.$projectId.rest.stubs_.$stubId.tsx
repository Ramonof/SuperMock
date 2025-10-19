import ProjectName from "@/components/generic/ProjectName"
import RestStubFullForm from "@/components/stub/rest/RestStubFullForm"
import RestStubList from "@/components/stub/rest/RestStubList"
import StubsInfo from "@/components/stub/rest/StubsInfo"
import { Stack, Container, VStack, HStack, Grid, GridItem, Flex, Box } from "@chakra-ui/react"
import { createFileRoute } from "@tanstack/react-router"

export const Route = createFileRoute('/_authenticated/project_/$projectId/rest/stubs_/$stubId')({
  component: RestStubs,
})

function RestStubs() {
  const { projectId, stubId } = Route.useParams()
  
  return (
      <>
        <ProjectName ProjectId={projectId}/>
        <StubsInfo ProjectId={projectId}/>
        <Grid
          templateAreas={`"header header"
                          "nav main"
                          "nav footer"`}
          gridTemplateRows={'50px 1fr 30px'}
          gridTemplateColumns={'150px 1fr'}
          h='200px'
          gap='1'
          color='blackAlpha.700'
          fontWeight='bold'
        >
          <GridItem pl='2' bg='orange.300' area={'header'}>
            Header
          </GridItem>
          <GridItem pl='2' bg='pink.300' area={'nav'}>
            Nav
          </GridItem>
          <GridItem pl='2' area={'main'}>
            <Grid templateColumns="repeat(2, 1fr)" gap="6">
              <Box pr={10}
                borderRight={"1px"}
                borderColor={"gray.600"}>
                <RestStubList ProjectId={projectId}/>
              </Box>
              <Box>
                <RestStubFullForm ProjectId={projectId} StubId={stubId}/>
              </Box>
            </Grid>
          </GridItem>
          <GridItem pl='2' bg='blue.300' area={'footer'}>
            Footer
          </GridItem>
        </Grid>
      </>
  )
}

// function RestStubs() {
//   const { projectId } = Route.useParams()
  
//   return (
//       <>
//         <ProjectName ProjectId={projectId}/>
//         <StubsInfo ProjectId={projectId}/>
//         <Grid
//           templateAreas={`"header header"
//                           "nav main"
//                           "nav footer"`}
//           gridTemplateRows={'50px 1fr 30px'}
//           gridTemplateColumns={'150px 1fr'}
//           h='200px'
//           gap='1'
//           color='blackAlpha.700'
//           fontWeight='bold'
//         >
//           <GridItem pl='2' bg='orange.300' area={'header'}>
//             Header
//           </GridItem>
//           <GridItem pl='2' bg='pink.300' area={'nav'}>
//             Nav
//           </GridItem>
//           <GridItem pl='2' bg='green.300' area={'main'}>
//             <Grid templateColumns="repeat(2, 1fr)" gap="6">
//               <Box pr={10}
//                 borderRight={"1px"}
//                 borderColor={"gray.600"}>
//                 <RestStubList ProjectId={projectId}/>
//               </Box>
//               <Box>
//                 <RestStubFullForm ProjectId={projectId}/>
//               </Box>
//             </Grid>
//           </GridItem>
//           <GridItem pl='2' bg='blue.300' area={'footer'}>
//             Footer
//           </GridItem>
//         </Grid>
//       </>
//   )
// }