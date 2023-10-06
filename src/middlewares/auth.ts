import { getUserById } from "resolvers";
// @ts-ignore
export const isAuthenticated = async ({set, jwt, cookie}) =>{
        if (!cookie!.access_token) {
            set.status = 401;
            return {
                success: false,
                message: "Unauthorized",
                data: null,
            };
        }
        const { userId } = await jwt.verify(cookie!.access_token);
        if (!userId) {
            set.status = 401;
            return {
                success: false,
                message: "Unauthorized",
                data: null,
            };
        }

        const user = await getUserById(userId);
        if (user.length === 0) {
            set.status = 401;
            return {
                success: false,
                message: "Unauthorized",
                data: null,
            };
        }
        return {
            user,
        };
}
