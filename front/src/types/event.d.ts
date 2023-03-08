export type Event = {
    ID: number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt: Date;
    Title: string;
    Description: string;
    Location: string;
    StartTime: string;
    EndTime: string;
    Owner: string;
    Subscribers: string[];
    MaxSubscribers: number;
}