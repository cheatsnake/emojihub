import { Module } from "@nestjs/common";
import { MongooseModule } from "@nestjs/mongoose";
import { AppController } from "./app.controller";
import { AppService } from "./app.service";
import { EmojiModule } from "./emoji/emoji.module";

@Module({
    imports: [
        EmojiModule,
        MongooseModule.forRoot(
            "mongodb+srv://admin:ANCp6ArZZ!5JSxx@cluster0.wvpsn.mongodb.net/EmojiDB?"
        ),
    ],
    controllers: [AppController],
    providers: [AppService],
})
export class AppModule {}
