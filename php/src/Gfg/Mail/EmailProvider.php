<?php
namespace Gfg\Mail;

class EmailProvider
{
    const FROM_EMAIL = 'php.coding.challenge@global-fashion-group.com';

    /** @var string */
    private $fileToLog;

    public function __construct(string $fileToLog)
    {
        $this->fileToLog = $fileToLog;
    }

    public function sendStockChangedEmail(string $productName, int $previousStock, int $newStock, string $email)
    {
        $subject = 'Stock updated';

        $message = sprintf(
            'The stock of your product "%s" has been modified from %d to %d',
            $productName,
            $previousStock,
            $newStock
        );

        $this->sendMail(static::FROM_EMAIL, $email, $subject, $message);
    }

    private function sendMail(string $from, string $to, string $subject, string $message)
    {
        $log = sprintf(
            "Email sent from %s to %s. Subject: %s; Message: %s\n",
            $from,
            $to,
            $subject,
            $message
        );

        // Lets imagine this is actually calling some external mailing service
        file_put_contents($this->fileToLog, $log, FILE_APPEND);
    }
}
