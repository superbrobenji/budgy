"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
import { Stack, StackProps } from 'aws-cdk-lib'
import { Construct } from 'constructs'
import { AttributeType, BillingMode, Table } from 'aws-cdk-lib/aws-dynamodb';

export class DatabaseStack extends Stack {
    public readonly tables = {}
    constructor(scope: Construct, id: string, props: StackProps) {
        super(scope, id, props);
        //DynamoDB tables
        const categoriesTable = new Table(this, 'dynamodbCategoriesStack', {
            partitionKey: { name: 'Id', type: AttributeType.STRING },
            billingMode: BillingMode.PAY_PER_REQUEST, // Use on-demand billing
            tableName: 'categories'
        });
        const itemsTable = new Table(this, 'dynamodbItemsStack', {
            partitionKey: { name: 'Id', type: AttributeType.STRING },
            billingMode: BillingMode.PAY_PER_REQUEST, // Use on-demand billing
            tableName: 'items'
        });
        const transactionsTable = new Table(this, 'dynamodbTransactionsStack', {
            partitionKey: { name: 'Id', type: AttributeType.STRING },
            billingMode: BillingMode.PAY_PER_REQUEST, // Use on-demand billing
            tableName: 'transactions'
        });
        this.tables = {
            categoriesTable,
            itemsTable,
            transactionsTable
        }
    }
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiY2RrLXN0YWNrLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiY2RrLXN0YWNrLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7OztBQUFBLDJDQUFrRTtBQUNsRSxpQ0FBZ0M7QUFFaEMscURBQXFEO0FBQ3JELG1EQUFtRDtBQUNuRCw4RUFBMEU7QUFDMUUsb0VBQW1FO0FBQ25FLDZCQUE4QjtBQUM5QixtRUFBNkY7QUFFN0YsNkZBQWtGO0FBQ2xGLDZDQUE2QztBQUM3QywrREFBK0Q7QUFDL0QsZ0VBQWdFO0FBQ2hFLDhCQUE4QjtBQUM5QixtRkFBbUY7QUFFbkYsTUFBYSxVQUFXLFNBQVEsWUFBSztJQUNqQyxZQUFZLEtBQWdCLEVBQUUsRUFBVSxFQUFFLEtBQWtCO1FBQ3hELEtBQUssQ0FBQyxLQUFLLEVBQUUsRUFBRSxFQUFFLEtBQUssQ0FBQyxDQUFDO1FBQ3hCLE1BQU0sQ0FBQyxNQUFNLEVBQUUsQ0FBQTtRQUVmLGlCQUFpQjtRQUNqQixNQUFNLGVBQWUsR0FBRyxJQUFJLFFBQVEsQ0FBQyxLQUFLLENBQUMsSUFBSSxFQUFFLHlCQUF5QixFQUFFO1lBQ3hFLFlBQVksRUFBRSxFQUFFLElBQUksRUFBRSxJQUFJLEVBQUUsSUFBSSxFQUFFLFFBQVEsQ0FBQyxhQUFhLENBQUMsTUFBTSxFQUFFO1lBQ2pFLFdBQVcsRUFBRSxRQUFRLENBQUMsV0FBVyxDQUFDLGVBQWUsRUFBRSx3QkFBd0I7WUFDM0UsU0FBUyxFQUFFLFlBQVk7U0FDMUIsQ0FBQyxDQUFDO1FBRUgsTUFBTSxVQUFVLEdBQUcsSUFBSSxRQUFRLENBQUMsS0FBSyxDQUFDLElBQUksRUFBRSxvQkFBb0IsRUFBRTtZQUM5RCxZQUFZLEVBQUUsRUFBRSxJQUFJLEVBQUUsSUFBSSxFQUFFLElBQUksRUFBRSxRQUFRLENBQUMsYUFBYSxDQUFDLE1BQU0sRUFBRTtZQUNqRSxXQUFXLEVBQUUsUUFBUSxDQUFDLFdBQVcsQ0FBQyxlQUFlLEVBQUUsd0JBQXdCO1lBQzNFLFNBQVMsRUFBRSxP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RUFBRSxXQUFXLEdBQUcsS0FBSzt3QkFDekIsT0FBTyxFQUFFLENBQUMsTUFBTSxDQUFDO3dCQUNqQixXQUFXLEVBQUUsZ0JBQWdCO3FCQUNoQyxDQUFDLENBQUM7Z0JBQ1AsQ0FBQztZQUNMLENBQUM7UUFDTCxDQUFDO1FBR0QsdUJBQXVCO1FBQ3ZCLG9FQUFvRTtRQUNwRSxpRUFBaUU7UUFDakUsNENBQTRDO1FBQzVDLHNDQUFzQztRQUN0QyxhQUFhO1FBQ2IsRUFBRTtRQUNGLGtEQUFrRDtRQUNsRCx1REFBdUQ7UUFDdkQsYUFBYTtRQUNiLEVBQUU7UUFDRiw4REFBOEQ7UUFDOUQsc0JBQXNCO1FBQ3RCLGFBQWE7UUFDYixFQUFFO1FBQ0Ysc0VBQXNFO1FBQ3RFLDRHQUE0RztRQUM1RywyQ0FBMkM7UUFDM0MsMEZBQTBGO1FBQzFGLGFBQWE7UUFDYixFQUFFO1FBQ0Ysd0VBQXdFO1FBQ3hFLGlEQUFpRDtRQUNqRCwyQkFBMkI7UUFDM0Isc0NBQXNDO1FBQ3RDLG9FQUFvRTtRQUNwRSxlQUFlO1FBQ2YsYUFBYTtRQUNiLEVBQUU7UUFDRiwyREFBMkQ7UUFDM0Qsc0NBQXNDO1FBQ3RDLHFHQUFxRztRQUNyRyxhQUFhO1FBQ2IsRUFBRTtRQUNGLHFGQUFxRjtRQUNyRixtQ0FBbUM7UUFDbkMsNENBQTRDO1FBQzVDLHlDQUF5QztRQUN6QyxzRUFBc0U7UUFDdEUsb0VBQW9FO1FBQ3BFLDRDQUE0QztRQUM1QywyREFBMkQ7UUFDM0QseUtBQXlLO1FBQ3pLLGFBQWE7UUFDYixFQUFFO1FBQ0YsdUNBQXVDO1FBQ3ZDLG1DQUFtQztRQUNuQyxtRkFBbUY7UUFDbkYsd0RBQXdEO1FBQ3hELFlBQVk7UUFDWixFQUFFO1FBQ0Ysb0RBQW9EO1FBQ3BELHdFQUF3RTtRQUN4RSw2QkFBNkI7UUFDN0IsWUFBWTtJQUNoQixDQUFDO0NBQ0o7QUFqSUQsZ0NBaUlDIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgQ2ZuUmVzb3VyY2UsIFN0YWNrLCBTdGFja1Byb3BzIH0gZnJvbSAnYXdzLWNkay1saWIvY29yZSc7XG5pbXBvcnQgKiBhcyBkb3RlbnYgZnJvbSAnZG90ZW52J1xuaW1wb3J0IHsgQ29uc3RydWN0IH0gZnJvbSAnY29uc3RydWN0cyc7XG5pbXBvcnQgKiBhcyBkeW5hbW9kYiBmcm9tICdhd3MtY2RrLWxpYi9hd3MtZHluYW1vZGInO1xuaW1wb3J0ICogYXMgZ28gZnJvbSAnQGF3cy1jZGsvYXdzLWxhbWJkYS1nby1hbHBoYSc7XG5pbXBvcnQgZ2V0RmlsZU5hbWVzQW5kRGlyZWN0b3JpZXMgZnJvbSAnLi4vdXRpbHMvZ2VuZXJhdGVMYW1iZGFGdW5jdGlvbnMnO1xuaW1wb3J0ICogYXMgcm91dGVEZWZzIGZyb20gJy4uLy4uLy4uL3RyYW5zcG9ydC9odHRwL3JvdXRlRGVmcy5qc29uJ1xuaW1wb3J0IHBhdGggPSByZXF1aXJlKCdwYXRoJyk7XG5pbXBvcnQgeyBDZm5JbnRlZ3JhdGlvbiwgQ2ZuUm91dGUsIEh0dHBBcGksIEh0dHBNZXRob2QgfSBmcm9tICdhd3MtY2RrLWxpYi9hd3MtYXBpZ2F0ZXdheXYyJztcbmltcG9ydCAqIGFzIGVjMiBmcm9tICdhd3MtY2RrLWxpYi9hd3MtZWMyJztcbmltcG9ydCB7IEh0dHBMYW1iZGFJbnRlZ3JhdGlvbiB9IGZyb20gJ2F3cy1jZGstbGliL2F3cy1hcGlnYXRld2F5djItaW50ZWdyYXRpb25zJztcbi8vaW1wb3J0ICogYXMgZWNzIGZyb20gJ2F3cy1jZGstbGliL2F3cy1lY3MnO1xuLy9pbXBvcnQgKiBhcyBlY3NfcGF0dGVybnMgZnJvbSAnYXdzLWNkay1saWIvYXdzLWVjcy1wYXR0ZXJucyc7XG4vL2ltcG9ydCB7IERvY2tlckltYWdlQXNzZXQgfSBmcm9tICdhd3MtY2RrLWxpYi9hd3MtZWNyLWFzc2V0cyc7XG4vL2ltcG9ydCB7IGpvaW4gfSBmcm9tICdwYXRoJztcbi8vaW1wb3J0IHsgQ2ZuSW50ZWdyYXRpb24sIENmblJvdXRlLCBIdHRwQXBpIH0gZnJvbSAnYXdzLWNkay1saWIvYXdzLWFwaWdhdGV3YXl2Mic7XG5cbmV4cG9ydCBjbGFzcyBCdWRneVN0YWNrIGV4dGVuZHMgU3RhY2sge1xuICAgIGNvbnN0cnVjdG9yKHNjb3BlOiBDb25zdHJ1Y3QsIGlkOiBzdHJpbmcsIHByb3BzPzogU3RhY2tQcm9wcykge1xuICAgICAgICBzdXBlcihzY29wZSwgaWQsIHByb3BzKTtcbiAgICAgICAgZG90ZW52LmNvbmZpZygpXG5cbiAgICAgICAgLy9EeW5hbW9EQiB0YWJsZXNcbiAgICAgICAgY29uc3QgY2F0ZWdvcmllc1RhYmxlID0gbmV3IGR5bmFtb2RiLlRhYmxlKHRoaXMsICdkeW5hbW9kYkNhdGVnb3JpZXNTdGFjaycsIHtcbiAgICAgICAgICAgIHBhcnRpdGlvbktleTogeyBuYW1lOiAnSWQnLCB0eXBlOiBkeW5hbW9kYi5BdHRyaWJ1dGVUeXBlLlNUUklORyB9LFxuICAgICAgICAgICAgYmlsbGluZ01vZGU6IGR5bmFtb2RiLkJpbGxpbmdNb2RlLlBBWV9QRVJfUkVRVUVTVCwgLy8gVXNlIG9uLWRlbWFuZCBiaWxsaW5nXG4gICAgICAgICAgICB0YWJsZU5hbWU6ICdjYXRlZ29yaWVzJ1xuICAgICAgICB9KTtcblxuICAgICAgICBjb25zdCBpdGVtc1RhYmxlID0gbmV3IGR5bmFtb2RiLlRhYmxlKHRoaXMsICdkeW5hbW9kYkl0ZW1zU3RhY2snLCB7XG4gICAgICAgICAgICBwYXJ0aXRpb25LZXk6IHsgbmFtZTogJ0lkJywgdHlwZTogZHluYW1vZGIuQXR0cmlidXRlVHlwZS5TVFJJTkcgfSxcbiAgICAgICAgICAgIGJpbGxpbmdNb2RlOiBkeW5hbW9kYi5CaWxsaW5nTW9kZS5QQVlfUEVSX1JFUVVFU1QsIC8vIFVzZSBvbi1kZW1hbmQgYmlsbGluZ1xuICAgICAgICAgICAgdGFibGVOYW1lOiAnaXRlbXMnXG4gICAgICAgIH0pO1xuICAgICAgICBjb25zdCB0cmFuc2FjdGlvbnNUYWJsZSA9IG5ldyBkeW5hbW9kYi5UYWJsZSh0aGlzLCAnZHluYW1vZGJUcmFuc2FjdGlvbnNTdGFjaycsIHtcbiAgICAgICAgICAgIHBhcnRpdGlvbktleTogeyBuYW1lOiAnSWQnLCB0eXBlOiBkeW5hbW9kYi5BdHRyaWJ1dGVUeXBlLlNUUklORyB9LFxuICAgICAgICAgICAgYmlsbGluZ01vZGU6IGR5bmFtb2RiLkJpbGxpbmdNb2RlLlBBWV9QRVJfUkVRVUVTVCwgLy8gVXNlIG9uLWRlbWFuZCBiaWxsaW5nXG4gICAgICAgICAgICB0YWJsZU5hbWU6ICd0cmFuc2FjdGlvbnMnXG4gICAgICAgIH0pO1xuXG4gICAgICAgIC8vIEFQSSBHYXRld2F5XG4gICAgICAgIGNvbnN0IGh0dHBBcGkgPSBuZXcgSHR0cEFwaSh0aGlzLCAnQnVkZ3lIdHRwQXBpJywge1xuICAgICAgICAgICAgYXBpTmFtZTogJ0J1ZGd5QXBpJyxcbiAgICAgICAgICAgIGRlc2NyaXB0aW9uOiAnVGhpcyBpcyB0aGUgQnVkZ3kgQVBJJyxcbiAgICAgICAgfSk7XG5cbiAgICAgICAgLy8gTGFtYmRhIGZ1bmN0aW9uc1xuICAgICAgICBjb25zdCBtZXRob2RGYWN0b3J5ID0gKG1ldGhvZDogc3RyaW5nKSA9PiB7XG4gICAgICAgICAgICBzd2l0Y2ggKG1ldGhvZCkge1xuICAgICAgICAgICAgICAgIGNhc2UgJ0dFVCc6XG4gICAgICAgICAgICAgICAgICAgIHJldHVybiBIdHRwTWV0aG9kLkdFVDtcbiAgICAgICAgICAgICAgICBjYXNlICdQT1NUJzpcbiAgICAgICAgICAgICAgICAgICAgcmV0dXJuIEh0dHBNZXRob2QuUE9TVDtcbiAgICAgICAgICAgICAgICBjYXNlICdQVVQnOlxuICAgICAgICAgICAgICAgICAgICByZXR1cm4gSHR0cE1ldGhvZC5QVVQ7XG4gICAgICAgICAgICAgICAgY2FzZSAnREVMRVRFJzpcbiAgICAgICAgICAgICAgICAgICAgcmV0dXJuIEh0dHBNZXRob2QuREVMRVRFO1xuICAgICAgICAgICAgICAgIGRlZmF1bHQ6XG4gICAgICAgICAgICAgICAgICAgIHJldHVybiBIdHRwTWV0aG9kLkFOWTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICAvL2xhbWJkYSBmdW5jdGlvbnMgZmFjdG9yeVxuICAgICAgICBjb25zdCBiYXNlQXBpUGF0aCA9IHByb2Nlc3MuZW52LlJPVVRFX0FQSV9FTkRQT0lOVCArIFwiL1wiICsgcHJvY2Vzcy5lbnYuVkVSU0lPTjtcbiAgICAgICAgY29uc3QgbGFtYmRhUGF0aCA9IHBhdGguam9pbihfX2Rpcm5hbWUsICcuLicsICcuLicsICcuLicsICd0cmFuc3BvcnQnLCAnaHR0cCcsICdsYW1iZGFIYW5kbGVycycpO1xuICAgICAgICBjb25zdCB7IGZpbGVOYW1lcywgZGlyZWN0b3JpZXMgfSA9IGdldEZpbGVOYW1lc0FuZERpcmVjdG9yaWVzKGxhbWJkYVBhdGgpO1xuICAgICAgICBmb3IgKGxldCBpID0gMDsgaSA8IGZpbGVOYW1lcy5sZW5ndGg7IGkrKykge1xuICAgICAgICAgICAgY29uc3QgbGFtYmRhRnVuYyA9IG5ldyBnby5Hb0Z1bmN0aW9uKHRoaXMsIGZpbGVOYW1lc1tpXSwge1xuICAgICAgICAgICAgICAgIGVudHJ5OiBwYXRoLmpvaW4oZGlyZWN0b3JpZXNbaV0pLFxuICAgICAgICAgICAgfSk7XG4gICAgICAgICAgICAvL1RPRE8gY3JlYXRlIGR5bmFtaWMgdGFibGUgYWNjZXNzIGZvciBlYWNoIGxhbWJkYSBmdW5jdGlvblxuICAgICAgICAgICAgaXRlbXNUYWJsZS5ncmFudFJlYWRXcml0ZURhdGEobGFtYmRhRnVuYyk7XG4gICAgICAgICAgICB0cmFuc2FjdGlvbnNUYWJsZS5ncmFudFJlYWRXcml0ZURhdGEobGFtYmRhRnVuYyk7XG4gICAgICAgICAgICBjYXRlZ29yaWVzVGFibGUuZ3JhbnRSZWFkV3JpdGVEYXRhKGxhbWJkYUZ1bmMpO1xuICAgICAgICAgICAgLy8gQVBJIEdhdGV3YXkgcm91dGUgZmFjdG9yeVxuICAgICAgICAgICAgZm9yIChjb25zdCByb3V0ZURlZiBvZiByb3V0ZURlZnMpIHtcbiAgICAgICAgICAgICAgICBjb25zdCBzZXJ2aWNlTmFtZSA9IHBhdGguYmFzZW5hbWUoZGlyZWN0b3JpZXNbaV0pO1xuICAgICAgICAgICAgICAgIGlmIChyb3V0ZURlZi5oYW5kbGVyID09PSBzZXJ2aWNlTmFtZSkge1xuICAgICAgICAgICAgICAgICAgICBjb25zdCByb3V0ZSA9IHJvdXRlRGVmLnJvdXRlO1xuICAgICAgICAgICAgICAgICAgICBjb25zdCBtZXRob2QgPSBtZXRob2RGYWN0b3J5KHJvdXRlRGVmLm1ldGhvZClcbiAgICAgICAgICAgICAgICAgICAgY29uc3QgbGFtYmRhSW5lZ3JhdGlvbiA9IG5ldyBIdHRwTGFtYmRhSW50ZWdyYXRpb24oZmlsZU5hbWVzW2ldLCBsYW1iZGFGdW5jKTtcbiAgICAgICAgICAgICAgICAgICAgaHR0cEFwaS5hZGRSb3V0ZXMoe1xuICAgICAgICAgICAgICAgICAgICAgICAgcGF0aDogYmFzZUFwaVBhdGggKyByb3V0ZSxcbiAgICAgICAgICAgICAgICAgICAgICAgIG1ldGhvZHM6IFttZXRob2RdLFxuICAgICAgICAgICAgICAgICAgICAgICAgaW50ZWdyYXRpb246IGxhbWJkYUluZWdyYXRpb24sXG4gICAgICAgICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgICAgIH1cbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuXG5cbiAgICAgICAgLy8gICAgICAgIC8vRUNTIEZhcmdhdGVcbiAgICAgICAgLy8gICAgICAgIGNvbnN0IGltYWdlID0gbmV3IERvY2tlckltYWdlQXNzZXQodGhpcywgXCJCYWNrZW5kSW1hZ2VcIiwge1xuICAgICAgICAvLyAgICAgICAgICAgIGRpcmVjdG9yeTogam9pbihfX2Rpcm5hbWUsIFwiLi5cIiwgXCIuLlwiLCBcIi4uXCIsIFwiLi5cIiksXG4gICAgICAgIC8vICAgICAgICAgICAgZmlsZTogXCJEb2NrZXJmaWxlLm11bHRpc3RhZ2VcIixcbiAgICAgICAgLy8gICAgICAgICAgICB0YXJnZXQ6IFwicmVsZWFzZS1zdGFnZVwiLFxuICAgICAgICAvLyAgICAgICAgfSk7XG4gICAgICAgIC8vXG4gICAgICAgIC8vICAgICAgICBjb25zdCB2cGMgPSBuZXcgZWMyLlZwYyh0aGlzLCBcIk15VnBjXCIsIHtcbiAgICAgICAgLy8gICAgICAgICAgICBtYXhBenM6IDMgLy8gRGVmYXVsdCBpcyBhbGwgQVpzIGluIHJlZ2lvblxuICAgICAgICAvLyAgICAgICAgfSk7XG4gICAgICAgIC8vXG4gICAgICAgIC8vICAgICAgICBjb25zdCBjbHVzdGVyID0gbmV3IGVjcy5DbHVzdGVyKHRoaXMsIFwiTXlDbHVzdGVyXCIsIHtcbiAgICAgICAgLy8gICAgICAgICAgICB2cGM6IHZwY1xuICAgICAgICAvLyAgICAgICAgfSk7XG4gICAgICAgIC8vXG4gICAgICAgIC8vICAgICAgICAvLyBDcmVhdGUgYSBsb2FkLWJhbGFuY2VkIEZhcmdhdGUgc2VydmljZSBhbmQgbWFrZSBpdCBwdWJsaWNcbiAgICAgICAgLy8gICAgICAgIGNvbnN0IGZhcmdhdGUgPSBuZXcgZWNzX3BhdHRlcm5zLkFwcGxpY2F0aW9uTG9hZEJhbGFuY2VkRmFyZ2F0ZVNlcnZpY2UodGhpcywgXCJNeUZhcmdhdGVTZXJ2aWNlXCIsIHtcbiAgICAgICAgLy8gICAgICAgICAgICBjbHVzdGVyOiBjbHVzdGVyLCAvLyBSZXF1aXJlZFxuICAgICAgICAvLyAgICAgICAgICAgIHRhc2tJbWFnZU9wdGlvbnM6IHsgaW1hZ2U6IGVjcy5Db250YWluZXJJbWFnZS5mcm9tRG9ja2VySW1hZ2VBc3NldChpbWFnZSkgfSxcbiAgICAgICAgLy8gICAgICAgIH0pO1xuICAgICAgICAvL1xuICAgICAgICAvLyAgICAgICAgY29uc3QgaHR0cFZwY0xpbmsgPSBuZXcgY2RrLkNmblJlc291cmNlKHRoaXMsICdIdHRwVnBjTGluaycsIHtcbiAgICAgICAgLy8gICAgICAgICAgICB0eXBlOiAnQVdTOjpBcGlHYXRld2F5VjI6OlZwY0xpbmsnLFxuICAgICAgICAvLyAgICAgICAgICAgIHByb3BlcnRpZXM6IHtcbiAgICAgICAgLy8gICAgICAgICAgICAgICAgTmFtZTogJ1YyIFZQQyBMaW5rJyxcbiAgICAgICAgLy8gICAgICAgICAgICAgICAgU3VibmV0SWRzOiB2cGMucHJpdmF0ZVN1Ym5ldHMubWFwKG0gPT4gbS5zdWJuZXRJZClcbiAgICAgICAgLy8gICAgICAgICAgICB9XG4gICAgICAgIC8vICAgICAgICB9KTtcbiAgICAgICAgLy9cbiAgICAgICAgLy8gICAgICAgIGNvbnN0IGFwaSA9IG5ldyBIdHRwQXBpKHRoaXMsICdIdHRwQXBpR2F0ZXdheScsIHtcbiAgICAgICAgLy8gICAgICAgICAgICBhcGlOYW1lOiAnQXBpZ3dGYXJnYXRlJyxcbiAgICAgICAgLy8gICAgICAgICAgICBkZXNjcmlwdGlvbjogJ0ludGVncmF0aW9uIGJldHdlZW4gYXBpZ3cgYW5kIEFwcGxpY2F0aW9uIExvYWQtQmFsYW5jZWQgRmFyZ2F0ZSBTZXJ2aWNlJyxcbiAgICAgICAgLy8gICAgICAgIH0pO1xuICAgICAgICAvL1xuICAgICAgICAvLyAgICAgICAgY29uc3QgaW50ZWdyYXRpb24gPSBuZXcgQ2ZuSW50ZWdyYXRpb24odGhpcywgJ0h0dHBBcGlHYXRld2F5SW50ZWdyYXRpb24nLCB7XG4gICAgICAgIC8vICAgICAgICAgICAgYXBpSWQ6IGFwaS5odHRwQXBpSWQsXG4gICAgICAgIC8vICAgICAgICAgICAgY29ubmVjdGlvbklkOiBodHRwVnBjTGluay5yZWYsXG4gICAgICAgIC8vICAgICAgICAgICAgY29ubmVjdGlvblR5cGU6ICdWUENfTElOSycsXG4gICAgICAgIC8vICAgICAgICAgICAgZGVzY3JpcHRpb246ICdBUEkgSW50ZWdyYXRpb24gd2l0aCBBV1MgRmFyZ2F0ZSBTZXJ2aWNlJyxcbiAgICAgICAgLy8gICAgICAgICAgICBpbnRlZ3JhdGlvbk1ldGhvZDogJ0dFVCcsIC8vIGZvciBHRVQgYW5kIFBPU1QsIHVzZSBBTllcbiAgICAgICAgLy8gICAgICAgICAgICBpbnRlZ3JhdGlvblR5cGU6ICdIVFRQX1BST1hZJyxcbiAgICAgICAgLy8gICAgICAgICAgICBpbnRlZ3JhdGlvblVyaTogZmFyZ2F0ZS5saXN0ZW5lci5saXN0ZW5lckFybixcbiAgICAgICAgLy8gICAgICAgICAgICBwYXlsb2FkRm9ybWF0VmVyc2lvbjogJzEuMCcsIC8vIHN1cHBvcnRlZCB2YWx1ZXMgZm9yIExhbWJkYSBwcm94eSBpbnRlZ3JhdGlvbnMgYXJlIDEuMCBhbmQgMi4wLiBGb3IgYWxsIG90aGVyIGludGVncmF0aW9ucywgMS4wIGlzIHRoZSBvbmx5IHN1cHBvcnRlZCB2YWx1ZVxuICAgICAgICAvLyAgICAgICAgfSk7XG4gICAgICAgIC8vXG4gICAgICAgIC8vICAgICAgICBuZXcgQ2ZuUm91dGUodGhpcywgJ1JvdXRlJywge1xuICAgICAgICAvLyAgICAgICAgICAgIGFwaUlkOiBhcGkuaHR0cEFwaUlkLFxuICAgICAgICAvLyAgICAgICAgICAgIHJvdXRlS2V5OiAnR0VUIC8nLCAgLy8gZm9yIHNvbWV0aGluZyBtb3JlIGdlbmVyYWwgdXNlICdBTlkgL3twcm94eSt9J1xuICAgICAgICAvLyAgICAgICAgICAgIHRhcmdldDogYGludGVncmF0aW9ucy8ke2ludGVncmF0aW9uLnJlZn1gLFxuICAgICAgICAvLyAgICAgICAgfSlcbiAgICAgICAgLy9cbiAgICAgICAgLy8gICAgICAgIG5ldyBjZGsuQ2ZuT3V0cHV0KHRoaXMsICdBUElHYXRld2F5VXJsJywge1xuICAgICAgICAvLyAgICAgICAgICAgIGRlc2NyaXB0aW9uOiAnQVBJIEdhdGV3YXkgVVJMIHRvIGFjY2VzcyB0aGUgR0VUIGVuZHBvaW50JyxcbiAgICAgICAgLy8gICAgICAgICAgICB2YWx1ZTogYXBpLnVybCFcbiAgICAgICAgLy8gICAgICAgIH0pXG4gICAgfVxufVxuXG4iXX0=
