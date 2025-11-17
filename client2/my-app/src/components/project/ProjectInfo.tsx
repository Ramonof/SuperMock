import { Badge, Box, Flex, Link, Spinner, Stack, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { Link as TanstackLink } from "@tanstack/react-router";
import getColor from "@/utils/color";

const ProjectInfo = ({ ProjectId }: { ProjectId: string }) => {

    return (
        <>
            <Stack gap={3}>
                <Text
                    fontSize={"4xl"}
                    textTransform={"uppercase"}
                    fontWeight={"bold"}
                    textAlign={"center"}
                    my={2}
                    bgGradient='linear(to-l, #0bf827ff, #4000ffff)'
                    bgClip='text'
                >
                    Info About Project {ProjectId}
                </Text>
                <Flex gap={2} alignItems={"center"}>
                    <Flex
                        flex={1}
                        alignItems={"center"}
                        border={"1px"}
                        borderColor={"gray.600"}
                        p={2}
                        borderRadius={"lg"}
                        justifyContent={"space-between"}
                    >
                        <Link as={TanstackLink}
                            to={`/project/${ProjectId}/rest/stubs/refactorme`}
                            color={getColor()}
                            variant="underline"
                            colorPalette="teal"
                        >
                            {"Rest Stubs"}
                        </Link>{" "}
                    </Flex>
                    <Flex
                        flex={1}
                        alignItems={"center"}
                        border={"1px"}
                        borderColor={"gray.600"}
                        p={2}
                        borderRadius={"lg"}
                        justifyContent={"space-between"}
                    >
                        <Link as={TanstackLink}
                            to={`/project/${ProjectId}/rest/stubs/refactorme`}
                            color={getColor()}
                            variant="underline"
                            colorPalette="teal"
                        >
                            {"GRPC Stubs"}
                        </Link>{" "}
                    </Flex>
                    <Flex
                        flex={1}
                        alignItems={"center"}
                        border={"1px"}
                        borderColor={"gray.600"}
                        p={2}
                        borderRadius={"lg"}
                        justifyContent={"space-between"}
                    >
                        <Link as={TanstackLink}
                            to={`/project/${ProjectId}/rest/stubs/refactorme`}
                            color={getColor()}
                            variant="underline"
                            colorPalette="teal"
                        >
                            {"Kafka Stubs"}
                        </Link>{" "}
                    </Flex>
                </Flex>
                <Text color={getColor()}>User role: Project Admin</Text>
                <Text color={getColor()}>Project token: abcdefg16</Text>
            </Stack>
        </>
    )
};

export default ProjectInfo;