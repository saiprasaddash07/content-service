CREATE TABLE content (
    contentId int AUTO_INCREMENT PRIMARY KEY,
    title varchar(100) NOT NULL,
    story longtext NOT NULL,
    userId int NOT NULL,
    isDeleted varchar(10) NOT NULL DEFAULT 'false',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updatedAt datetime DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);
CREATE INDEX `contentId` ON content (contentId);
CREATE INDEX `userId` ON content (userId);
CREATE INDEX `title` ON content (title);
CREATE INDEX `createdAt` ON content (createdAt);
CREATE INDEX `isDeleted` ON content (isDeleted);