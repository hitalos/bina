angular.module('bina', [])
    .controller('buscaController', function($scope, $http){
        $http.get('/contatos/json').success(function(data){
            $scope.contatos = data;
        });

        $('#search').focus();
    })
    .directive('cardsCounter', function(){
        return {
            restrict: 'E',
            templateUrl: 'templates/cards-counter.html'
        };
    })
    .directive('card', function(){
        return {
            restrict: 'E',
            templateUrl: 'templates/card.html'
        };
    });
