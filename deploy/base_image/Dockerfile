FROM centos:7
ADD ./google-chrome-stable_current_x86_64.rpm /google-chrome-stable_current_x86_64.rpm
RUN yum -y install fontconfig
RUN yum -y localinstall /google-chrome-stable_current_x86_64.rpm
ADD WeiRuanYaHei.ttf /usr/share/fonts
RUN fc-cache -vf