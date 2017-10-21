package LWP::Debug::TraceHTTP;
$LWP::Debug::TraceHTTP::VERSION = '6.27';
# Just call:
#
#   require LWP::Debug::TraceHTTP;
#   LWP::Protocol::implementor('http', 'LWP::Debug::TraceHTTP');
#
# to use this module to trace all calls to the HTTP socket object in
# programs that use LWP.

use strict;
use base 'LWP::Protocol::http';

package # hide from PAUSE
    LWP::Debug::TraceHTTP::Socket;

use Data::Dump 1.13;
use Data::Dump::Trace qw(autowrap mcall);

autowrap("LWP::Protocol::http::Socket" => "sock");

sub new {
    my $class = shift;
    return mcall("LWP::Protocol::http::Socket" => "new", undef, @_);
}

1;
