import { Test, TestingModule } from "@nestjs/testing";
import { INestApplication } from "@nestjs/common";
import * as request from "supertest";
import { AppModule } from "./../src/app.module";
import { disconnect } from "mongoose";

describe("EmojiController (e2e)", () => {
    let app: INestApplication;

    beforeEach(async () => {
        const moduleFixture: TestingModule = await Test.createTestingModule({
            imports: [AppModule],
        }).compile();

        app = moduleFixture.createNestApplication();
        await app.init();
    });

    it("/api/all (GET)", async () => {
        return request(app.getHttpServer())
            .get("/api/all")
            .expect(200)
            .then(({ body }: request.Response) => {
                expect(body.length > 0);
            });
    });

    it("/api/all/category_smileys_and_people (GET)", async () => {
        return request(app.getHttpServer())
            .get("/api/all/category_smileys_and_people")
            .expect(200)
            .then(({ body }: request.Response) => {
                expect(body.length > 0);
            });
    });

    it("/api/all/group_face_positive (GET)", async () => {
        return request(app.getHttpServer())
            .get("/api/all/group_face_positive")
            .expect(200)
            .then(({ body }: request.Response) => {
                expect(body.length > 0);
            });
    });

    it("/api/random (GET)", async () => {
        return request(app.getHttpServer())
            .get("/api/random")
            .expect(200)
            .then(({ body }: request.Response) => {
                expect(body).toMatchObject({});
            });
    });

    it("/api/random/category_smileys_and_people (GET)", async () => {
        return request(app.getHttpServer())
            .get("/api/random/category_smileys_and_people")
            .expect(200)
            .then(({ body }: request.Response) => {
                expect(body).toMatchObject({});
            });
    });

    it("/api/random/group_face_positive (GET)", async () => {
        return request(app.getHttpServer())
            .get("/api/random/group_face_positive")
            .expect(200)
            .then(({ body }: request.Response) => {
                expect(body).toMatchObject({});
            });
    });

    afterEach(async () => {
        disconnect();
        await new Promise<void>((resolve) => setTimeout(() => resolve(), 500));
    });
});
