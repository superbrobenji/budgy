import { Elysia } from "elysia";
import modules from "modules";
import { cookie } from "@elysiajs/cookie";
import { jwt } from "@elysiajs/jwt";

const app = new Elysia();

app.get("/", () => "Hello Elysia");

app.group("/api", app =>
    app.use(
        jwt({
            name: "jwt",
            secret: process.env.JWT_SECRET!,
        })
    )
        .use(cookie())
        .use(modules)
);

app.listen(3000, () => {
    console.log(
        `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
    );
})

