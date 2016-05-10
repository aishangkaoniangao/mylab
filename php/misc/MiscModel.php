<?php
class MiscModel
{
    /*
     * 获取今天的0点: Unix Timestamp
     */
    static public function getTodayZeroSec()
    {
        $now = time();
        $date = sprintf("%s-%s-%s", date("Y", $now), date("m", $now), date("d", $now));
        return strtotime($date);
    }

    /*
     * 获取昨天的0点:Unix Timestamp
     */
    static public function getYesterdayZeroSec()
    {
        return self::getTodayZeroSec() - 86400;
    }

    static public function mailDomainAccount($emails, $subject, $message)
    {
        $service = new EmailService();
        $service->setSubject($subject);
        $service->setFrom('noreply@alarm.360.cn');
        $service->setTo($emails);
        $service->setBody($message);
        return $service->deliverIn();
    }

    static public function mailOutAccount($emails, $subject, $message)
    {
        $service = new EmailService();
        $service->setSubject($subject);
        $service->setTo($emails);
        $service->setBody($message);
        return $service->deliverOut();
    }
    
    static public function getCurrentUri()
    {
        return sprintf("%s://%s%s",
            $_SERVER["HTTPS"] == "on" ? "https" : "http",
            $_SERVER["HTTP_HOST"],
            $_SERVER["REQUEST_URI"]
        );
    }
}
