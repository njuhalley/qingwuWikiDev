cls
@echo off
REM ---------******************--------------
set THISDATETIME=%DATE:~0,4%%DATE:~5,2%%DATE:~8,2%-%TIME:~0,2%:%TIME:~3,2%.%TIME:~6,2%  
REM echo %THISDATETIME%
REM ---------------git�ϴ�----------------
D:
cd /
cd Documents
cd Coding
cd Python
cd QWResearchWork

REM ---------*******************************�ָ���*****************************--------------
echo.&call
echo ---------*********�ָ���*********--------------
REM ---------------QWResearchWork/pyQwSciLib----------------
echo.&call
echo ------QWResearchWork/pyQwSciLib-------
echo.&call

cd pyQwSciLib

git pull

git add .
git commit -m "commit QWResearchWork/pyQwSciLib %THISDATETIME%"
git push -u origin master


REM ---------------QWResearchWork/wSpider----------------
echo.&call
echo ------QWResearchWork/wSpider-------
echo.&call

cd ..
cd wSpider

git pull

git add .
git commit -m "commit QWResearchWork/wSpider %THISDATETIME%"
git push -u origin master

REM ---------*******************************�ָ���*****************************--------------
if defined test (
echo ---------*********�ָ���*********--------------
echo.&call

REM ---------------/pyQwSciLib----------------
echo.&call
echo ------/pyQwSciLib-------
echo.&call

cd ..
cd ..
cd ..
cd ..
cd Programs
cd pyQwSciLib

git pull

REM git add .
REM git commit -m "commit /pyQwSciLib %THISDATETIME%"
REM git push -u origin master


REM ---------------/wSpider----------------
echo.&call
echo ------/wSpider-------
echo.&call

cd ..
cd wSpider

git pull

REM git add .
REM git commit -m "commit /wSpider %THISDATETIME%"
REM git push -u origin master
) else (
echo ---------*********�ָ���*********--------------
echo ���ٸ���pyQwSciLib��wSpider
cd ..
cd ..
cd ..
cd ..
cd Programs
cd wSpider
)

REM ---------*******************************�ָ���*****************************--------------
echo ---------*********�ָ���*********--------------
REM ---------------/HalleySite----------------
echo.&call
echo ------/HalleySite-------
echo.&call

cd ..
cd HalleySite

git pull

cd wSpider
git pull

REM ---------*******************************�ָ���*****************************--------------
echo ---------*********�ָ���*********--------------
echo.&call
REM ---------------QWResearchWork---------------
echo ------QWResearchWork-------
echo.&call

cd ..
cd ..
cd ..
cd Coding
cd Python
cd QWResearchWork

git add .
git commit -m "commit QWResearchWork %THISDATETIME%"
git push -u origin master


REM ---------*******************************�ָ���*****************************--------------
echo.&call
echo ---------*********�ָ���*********--------------
REM ---------------QWResearchWork/egovResearch----------------
echo.&call
echo ------QWResearchWork/egovResearch-------
echo.&call

cd egovResearch

git add .
git commit -m "commit QWResearchWork/egovResearch %THISDATETIME%"
git push

REM ---------*******************************�ָ���*****************************--------------
echo ---------*********�ָ���*********--------------
echo.&call
exit

