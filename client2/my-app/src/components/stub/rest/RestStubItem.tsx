import { BASE_URL } from "@/main";
import { Flex, Spinner, Link, Badge, Box } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { Link as TanstackLink } from "@tanstack/react-router";
import type { RestStub } from "./RestStubList";
import getColor from "@/utils/color";

const RestStubItem = ({ Project }: { Project: RestStub }) => {
    const queryClient = useQueryClient();

    const { mutate: updateProject, isPending: isUpdating } = useMutation({
        mutationKey: ["updateProject"],
        mutationFn: async () => {
            if (Project.completed) return alert("Project is already completed");
            try {
                const res = await fetch(BASE_URL + `/projects/${Project._id}`, {
                    method: "PATCH",
                });
                const data = await res.json();
                if (!res.ok) {
                    throw new Error(data.error || "Something went wrong");
                }
                return data;
            } catch (error) {
                console.log(error);
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ["projects"] });
        },
    });

    const { mutate: deleteProject, isPending: isDeleting } = useMutation({
        mutationKey: ["deleteProject"],
        mutationFn: async () => {
            try {
                const res = await fetch(BASE_URL + `/projects/${Project._id}`, {
                    method: "DELETE",
                });
                const data = await res.json();
                if (!res.ok) {
                    throw new Error(data.error || "Something went wrong");
                }
                return data;
            } catch (error) {
                console.log(error);
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ["projects"] });
        },
    });

    return (
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
                    to={`/project/${Project.project_id}/rest/stubs/${Project.id}`}
                    color={Project.completed ? "green.200" : getColor}
                    textDecoration={Project.completed ? "line-through" : "none"}
                    variant="underline"
                    // href={`/project/${Project.id}`}
                    colorPalette="teal"
                >
                    {Project.name}
                </Link>{" "}
            </Flex>
            <Flex gap={2} alignItems={"center"}>
                {/* <Box color={"green.500"} cursor={"pointer"} onClick={() => updateProject()}>
                    {!isUpdating && <FaCheckCircle size={20} />}
                    {isUpdating && <Spinner size={"sm"} />}
                </Box> */}
                <Box color={"red.500"} cursor={"pointer"} onClick={() => deleteProject()}>
                    {!isDeleting && <MdDelete size={25} />}
                    {isDeleting && <Spinner size={"sm"} />}
                </Box>
            </Flex>
        </Flex>
    );
};
export default RestStubItem;