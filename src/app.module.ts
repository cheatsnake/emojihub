import { Module } from "@nestjs/common";
import { ConfigModule } from "@nestjs/config";
import { MongooseModule } from "@nestjs/mongoose";
import { AppController } from "./app.controller";
import { AppService } from "./app.service";
import { EmojiModule } from "./emoji/emoji.module";
import { ServeStaticModule } from "@nestjs/serve-static";
import { join } from "path";

@Module({
    imports: [
        ConfigModule.forRoot(),
        EmojiModule,
        MongooseModule.forRoot(process.env.DB_URL),
        ServeStaticModule.forRoot({
            rootPath: join(__dirname, "..", "public"),
            exclude: ["/api*"],
        }),
    ],
    controllers: [AppController],
    providers: [AppService],
})
export class AppModule {}
